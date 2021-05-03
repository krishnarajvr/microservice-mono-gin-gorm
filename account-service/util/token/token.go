package token

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"micro/config"
	"micro/module/user/repo"
	"micro/util/cache"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type TokenConfig struct {
	TokenRepo repo.ITokenRepo
	Cache     *cache.RedisClient
	Config    *config.Conf
}

type Token struct {
	CacheToken          bool
	Valid               bool
	TokenRepo           repo.ITokenRepo
	Client              *cache.RedisClient
	AccessSecret        string
	RefreshSecret       string
	AdminAccessExpiry   int
	AdminRefreshExpiry  int
	ClientAccessExpiry  int
	ClientRefreshExpiry int
}

type TokenData struct {
	AccessUuid   string
	UserId       int64
	TenantId     int64
	ReferenceId  string
	Subject      string
	Type         string
	Roles        []string
	Admin        bool
	RefreshToken string
}

type AccessDetails struct {
	AccessUuid  string
	UserId      int64
	TenantId    int64
	ReferenceId string
	Type        string
	Subject     string
	Roles       []string
	Admin       bool
}

type TokenDetails struct {
	AccessToken     string
	RefreshToken    string
	AccessUuid      string
	RefreshUuid     string
	AtExpires       int64
	RtExpires       int64
	ExpireInSeconds int
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	AtExpires    string `json:"accessTokenExpiry"`
	RtExpires    string `json:"refreshTokenExpiry"`
}

func (t *TokenDetails) ToResponse() TokenResponse {
	utcAtExpires := time.Unix(t.AtExpires, 0)
	rfcAtExpires := utcAtExpires.Format(time.RFC3339)
	utcRtExpires := time.Unix(t.RtExpires, 0)
	rfcRtExpires := utcRtExpires.Format(time.RFC3339)
	return TokenResponse{
		TokenType:    "bearer",
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		AtExpires:    rfcAtExpires,
		RtExpires:    rfcRtExpires,
	}
}

type OauthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	IssuedAt     string `json:".issued"`
	ExpiredAt    string `json:".expires"`
	ExpiresIn    int    `json:"expires_in"`
}

func (t *TokenDetails) OAuthResponse() OauthTokenResponse {
	utcAtExpires := time.Unix(t.AtExpires, 0)
	rfcAtExpires := utcAtExpires.Format(time.RFC3339)
	issuedAt := time.Unix(0, 0).Format(time.RFC3339)

	return OauthTokenResponse{
		TokenType:    "bearer",
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		ExpiredAt:    rfcAtExpires,
		ExpiresIn:    t.ExpireInSeconds,
		IssuedAt:     issuedAt,
	}
}

func New(tokenConfig TokenConfig) *Token {
	var token = Token{}
	token.AccessSecret = tokenConfig.Config.Token.AccessSecret
	token.RefreshSecret = tokenConfig.Config.Token.RefreshSecret
	token.AdminAccessExpiry = tokenConfig.Config.Token.AdminAccessExpiry
	token.AdminRefreshExpiry = tokenConfig.Config.Token.AdminRefreshExpiry
	token.ClientAccessExpiry = tokenConfig.Config.Token.ClientAccessExpiry
	token.ClientRefreshExpiry = tokenConfig.Config.Token.ClientRefreshExpiry

	token.TokenRepo = tokenConfig.TokenRepo

	token.CacheToken = false
	token.Valid = true

	if len(token.AccessSecret) == 0 || len(token.RefreshSecret) == 0 {
		token.Valid = false
		return &token
	}

	if tokenConfig.Cache != nil {
		token.CacheToken = true
		token.Client = tokenConfig.Cache
	}

	return &token
}

func (t *Token) CreateToken(tokenData *TokenData) (*TokenDetails, error) {
	td := &TokenDetails{}
	//Set expiry

	if strings.ToLower(tokenData.Type) == "vendor" {
		td.AtExpires = time.Now().Add(time.Minute * time.Duration(t.ClientAccessExpiry)).Unix()
		td.RtExpires = time.Now().Add(time.Minute * time.Duration(t.ClientRefreshExpiry)).Unix()
		td.ExpireInSeconds = t.ClientAccessExpiry * 60

	} else {
		td.AtExpires = time.Now().Add(time.Minute * time.Duration(t.AdminAccessExpiry)).Unix()
		td.RtExpires = time.Now().Add(time.Minute * time.Duration(t.AdminRefreshExpiry)).Unix()
		td.ExpireInSeconds = t.AdminAccessExpiry * 60
	}

	td.AccessUuid = uuid.NewV4().String()
	td.RefreshUuid = td.AccessUuid + "++" + strconv.Itoa(int(tokenData.UserId))

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["iss"] = tokenData.TenantId
	atClaims["sub"] = tokenData.Subject
	atClaims["admin"] = tokenData.Admin
	atClaims["type"] = tokenData.Type
	atClaims["access_uuid"] = td.AccessUuid

	if tokenData.UserId != 0 {
		atClaims["user_id"] = tokenData.UserId
	}

	if len(tokenData.ReferenceId) > 0 {
		atClaims["ref_id"] = tokenData.ReferenceId
	}

	atClaims["roles"] = tokenData.Roles
	atClaims["exp"] = td.AtExpires

	fmt.Println("AtClaims:", atClaims)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(t.AccessSecret))

	if err != nil {
		return nil, err
	}

	if len(tokenData.RefreshToken) > 0 {
		td.RefreshToken = tokenData.RefreshToken
	} else {
		rtClaims := jwt.MapClaims{}
		rtClaims["refresh_uuid"] = td.RefreshUuid
		rtClaims["exp"] = td.RtExpires
		rtClaims["iss"] = tokenData.TenantId

		if tokenData.UserId != 0 {
			rtClaims["user_id"] = tokenData.UserId
		}

		if len(tokenData.ReferenceId) > 0 {
			rtClaims["ref_id"] = tokenData.ReferenceId
		}

		rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
		td.RefreshToken, err = rt.SignedString([]byte(t.RefreshSecret))

		if err != nil {
			return nil, err
		}
	}

	return td, nil
}

//CreateAuth - Save the token details in cache if enabled
func (t *Token) CreateAuth(userid int64, td *TokenDetails) error {
	if !t.CacheToken {
		return nil
	}

	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := t.Client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now))

	if errAccess != nil {
		return errAccess
	}

	errRefresh := t.Client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now))

	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

func (t *Token) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// VerifyToken - Parse, validate, and return a token.
// keyFunc will receive the parsed token and should return the key for validating.
func (t *Token) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := t.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(t.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (t *Token) TokenValid(r *http.Request) error {
	token, err := t.VerifyToken(r)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return err
	}

	return nil
}

func (t *Token) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := t.VerifyToken(r)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token 1")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Invalid token 2")
	}

	refId := ""
	tokenType := ""
	accessUuid := ""
	sub := ""

	if claims["ref_id"] != nil {
		refId = claims["ref_id"].(string)
	}

	if claims["type"] != nil {
		tokenType = claims["type"].(string)
	}

	if claims["access_uuid"] != nil {
		accessUuid = claims["access_uuid"].(string)
	}

	if claims["sub"] != nil {
		sub = claims["sub"].(string)
	}

	userId, _ := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	tenantId, err2 := strconv.ParseInt(fmt.Sprintf("%.f", claims["iss"]), 10, 64)

	roles, rolesOk := claims["roles"].([]interface{})

	if !rolesOk {
		return nil, errors.New("Invalid Roles")
	}

	if len(roles) == 0 {
		return nil, errors.New("Invalid Roles")
	}

	roleString := roles[0].(string)

	if err2 != nil {
		return nil, errors.New("Invalid Tenant info")
	}

	ad := &AccessDetails{
		AccessUuid:  accessUuid,
		UserId:      int64(userId),
		TenantId:    int64(tenantId),
		ReferenceId: refId,
		Type:        tokenType,
		Subject:     sub,
		Roles:       []string{roleString},
	}

	return ad, nil
}

func (t *Token) FetchAuth(authD *AccessDetails) (int64, error) {
	if !t.CacheToken {
		return authD.UserId, nil
	}

	userid, err := t.Client.Get(authD.AccessUuid)

	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseInt(userid, 10, 64)

	if authD.UserId != userID {
		return 0, errors.New("unauthorized")
	}

	return userID, nil
}

func (t *Token) DeleteAuth(givenUuid string) (int64, error) {
	if !t.CacheToken {
		return 1, nil
	}

	deleted, err := t.Client.Delete(givenUuid)

	if err != nil {
		return 0, err
	}

	return deleted, nil
}

func (t *Token) Logout(c *gin.Context) {
	metadata, err := t.ExtractTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	delErr := t.DeleteTokens(metadata)

	if delErr != nil {
		c.JSON(http.StatusUnauthorized, delErr.Error())
		return
	}

	c.JSON(http.StatusOK, "Successfully logged out")
}

func (t *Token) Refresh(refreshToken string, tokenType string, subject string) (*TokenDetails, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(t.RefreshSecret), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		fmt.Println("the error: ", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	refreshUuid, ok := claims["refresh_uuid"].(string)

	if !ok {
		return nil, errors.New("Invalid Token")
	}

	tenantId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["iss"]), 10, 64)

	if err != nil {
		return nil, errors.New("Invalid Token data")
	}

	//Delete the previous Refresh Token
	deleted, delErr := t.DeleteAuth(refreshUuid)

	if delErr != nil || deleted == 0 {
		return nil, errors.New("Token not found")
	}

	var tokenData TokenData

	if tokenType == "user" {
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)

		if err != nil {
			return nil, errors.New("Invalid user Id")
		}

		tokenData = TokenData{
			UserId:       userId,
			TenantId:     tenantId,
			Subject:      subject,
			Admin:        true,
			Type:         tokenType,
			Roles:        []string{"admin"},
			RefreshToken: refreshToken,
		}
	} else {
		if claims["ref_id"] == nil {
			return nil, errors.New("Invalid reference Id")
		}

		tokenData = TokenData{
			ReferenceId:  claims["ref_id"].(string),
			TenantId:     tenantId,
			Subject:      subject,
			Admin:        false,
			Type:         tokenType,
			Roles:        []string{"vendor"},
			RefreshToken: refreshToken,
		}
	}

	ts, createErr := t.CreateToken(&tokenData)

	if createErr != nil {
		return nil, errors.New("Token creation failed")
	}

	return ts, nil
}

func (t *Token) DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.AccessUuid, authD.UserId)
	//delete access token
	deletedAt, err := t.Client.Delete(authD.AccessUuid)

	if err != nil {
		return err
	}

	deletedRt, err := t.Client.Delete(refreshUuid)
	if err != nil {
		return err
	}

	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}

	return nil
}
