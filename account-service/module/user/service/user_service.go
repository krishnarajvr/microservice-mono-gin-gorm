package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	common "github.com/krishnarajvr/micro-common"

	"micro/module/user/model"
	"micro/module/user/repo"
	util "micro/util"
	"micro/util/password"
	"micro/util/token"

	"github.com/krishnarajvr/micro-common/locale"
)

type IUserService interface {
	List(tenantId int, page common.Pagination, filters model.UserFilterList) (model.UserDtos, *common.PageResult, error)
	Get(tenantId int, id int) (*model.UserDto, error)
	Add(form *model.UserForm) (*model.UserDto, error)
	Update(form *model.UserForm, id int) (*model.UserDto, error)
	Patch(form *model.UserPatchForm, id int) (*model.UserDto, error)
	Delete(id int) (*model.UserDto, error)
	Login(login *model.LoginForm) (*model.User, error)
	GetToken(user *model.User) (*token.TokenDetails, error)
	ParseUri(Uri string) (map[string]string, bool)
	CheckPermission(role string, module string, resource string, method string) (bool, error)
	RefreshToken(refreshToken string) (*token.TokenDetails, error)
	ExtractTokenMetadata(r *http.Request) (*token.AccessDetails, error)
}

type ServiceConfig struct {
	UserRepo repo.IUserRepo
	Lang     *locale.Locale
	Token    *token.Token
}

type Service struct {
	UserRepo repo.IUserRepo
	Lang     *locale.Locale
	Token    *token.Token
}

func NewService(c ServiceConfig) IUserService {
	return &Service{
		UserRepo: c.UserRepo,
		Lang:     c.Lang,
		Token:    c.Token,
	}
}

func (s *Service) List(tenantId int, page common.Pagination, filters model.UserFilterList) (model.UserDtos, *common.PageResult, error) {
	users, pageResult, err := s.UserRepo.List(tenantId, page, filters)

	if err != nil {
		return nil, nil, err
	}

	usersDto := users.ToDto()

	return usersDto, pageResult, err
}

func (s *Service) Get(tenantId int, id int) (*model.UserDto, error) {
	user, err := s.UserRepo.Get(tenantId, id)

	if err != nil {
		return nil, err
	}

	userDto := user.ToDto()

	return userDto, nil
}

func (s *Service) Add(form *model.UserForm) (*model.UserDto, error) {
	userModel, err := form.ToModel()
	//Todo - Get from token
	userModel.TenantId = 1
	userModel.Password = password.Encrypt(form.Password)

	user, err := s.UserRepo.Add(userModel)

	if err != nil {
		return nil, err
	}

	userDto := user.ToDto()

	return userDto, nil
}

func (s *Service) Login(form *model.LoginForm) (*model.User, error) {
	user, err := s.UserRepo.GetByEmail(form.Email)

	if err != nil {
		return nil, err
	}

	passwordEqual := password.Compare(user.Password, form.Password)

	if passwordEqual != true {
		return nil, errors.New("Password not match")
	}

	return user, nil
}

func (s *Service) Update(form *model.UserForm, id int) (*model.UserDto, error) {
	user, err := s.UserRepo.Update(form, id)

	if err != nil {
		return nil, err
	}

	userDto := user.ToDto()

	return userDto, nil
}

func (s *Service) Patch(form *model.UserPatchForm, id int) (*model.UserDto, error) {
	user, err := s.UserRepo.Patch(form, id)

	if err != nil {
		return nil, err
	}

	userDto := user.ToDto()

	return userDto, nil
}

func (s *Service) Delete(id int) (*model.UserDto, error) {
	user, err := s.UserRepo.Delete(id)

	if err != nil {
		return nil, err
	}

	userDto := user.ToDto()

	return userDto, nil
}

func (s *Service) ExtractTokenMetadata(r *http.Request) (*token.AccessDetails, error) {
	return s.Token.ExtractTokenMetadata(r)
}

func (s *Service) GetToken(user *model.User) (*token.TokenDetails, error) {
	tokenData := token.TokenData{
		UserId:   int64(user.ID),
		TenantId: int64(user.TenantId),
		Subject:  "CarePlan",
		Admin:    true,
		Type:     "user",
		Roles:    []string{"admin"},
	}

	token, _ := s.Token.CreateToken(&tokenData)

	return token, nil
}

func (s *Service) RefreshToken(refreshToken string) (*token.TokenDetails, error) {
	token, err := s.Token.Refresh(refreshToken, "user", "CarePlan")

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *Service) ParseUri(uri string) (map[string]string, bool) {
	uriWithQueryParts := strings.Split(uri, "?")
	uriString := uriWithQueryParts[0]
	uriParts := strings.Split(uriString, "/")
	partLength := len(uriParts)

	if partLength <= 3 {
		return nil, false
	}

	module := uriParts[1]
	resource := uriParts[3]

	if partLength >= 5 {
		resource = resource + "/:id"
	}

	if len(module) == 0 || len(resource) == 0 {
		return nil, false
	}

	return map[string]string{
		"module":   module,
		"resource": resource,
	}, true
}

func (s *Service) CheckPermission(role string, module string, resource string, method string) (bool, error) {
	rolePemissionJson := s.getPermissions()

	var rolePermissions map[string]map[string]map[string][]string

	if errJson := json.Unmarshal([]byte(rolePemissionJson), &rolePermissions); errJson != nil {
		return false, errJson
	}

	methods, ok := rolePermissions[strings.ToLower(role)][strings.ToLower(module)][resource]

	if !ok {
		return false, nil
	}

	if util.ArrayContains(methods, strings.ToLower(method)) {
		return true, nil
	}

	return false, nil
}

//getPermissions Define the role permissions for API
//Todo - Need to move this to database and get from inmemory cache
func (s *Service) getPermissions() string {

	//role->module->url->permissions[]
	return `{
		"admin": {
			"account": {
				"users"     			:["get","post"],
				"users/:id" 			:["get","post","patch"],
				"tenants"     			:["get","post"],
				"tenants/:id" 			:["get","post","patch"],
				"clientCredentials" 	:["get","post"],
				"clientCredentials/:id" :["get","post","patch"]
			},
			"product": {
				"products"     			:["get","post"],
				"products/:id" 			:["get","post","patch"],
				"productTypes"     		:["get","post"],
				"productTypes/:id" 		:["get","post","patch"],
				"categories" 			:["get","post"],
				"categories/:id" 		:["get","post","patch"]
			}	
									          
		},
		"user": {
			"site": {
				"products"     			:["get"],
				"products/:id" 			:["get"],
				"productTypes"     		:["get"],
				"categories" 			:["get"],
				"categories/:id" 		:["get"]
			}
		}
	}`
}
