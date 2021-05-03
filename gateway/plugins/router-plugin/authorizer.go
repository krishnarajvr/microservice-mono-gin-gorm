package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type TokenData struct {
	UserId      string
	TenantId    string
	ReferenceId string
	Subject     string
	Type        string
	Roles       []string
	Admin       bool
}

type AuthErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AuthData struct {
	Auth TokenData `json:"auth"`
}

type AuthResponse struct {
	Status    int           `json:"status"`
	Data      AuthData      `json:"data"`
	Error     AuthErrorData `json:"error"`
	RequestId string        `json:"requestId"`
}

func GetAuthorizeContext(r *http.Request) (*AuthResponse, error) {
	client := &http.Client{}

	authorizeURL := os.Getenv("PERMISSION_URL")

	if len(authorizeURL) == 0 {
		return nil, errors.New("Invalid URL")
	}

	data := url.Values{}
	data.Add("uri", r.RequestURI)
	data.Add("method", r.Method)

	authReq, err := http.NewRequest("POST", authorizeURL, bytes.NewBufferString(data.Encode()))
	authReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	authReq.Header.Set("Authorization", r.Header.Get("Authorization"))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	authResp, err := client.Do(authReq)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	apiData, err := ioutil.ReadAll(authResp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	authResp.Body.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var authResponse AuthResponse

	if errJson := json.Unmarshal(apiData, &authResponse); errJson != nil {
		return nil, errJson
	}

	return &authResponse, nil
}
