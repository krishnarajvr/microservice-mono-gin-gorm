package main

import (
	"net/http"
	"strings"
)

func GetPublicRoutes() map[string]interface{} {
	return map[string]interface{}{
		"/thirdparty/v1/token":       "true",
		"/account/v1/adminLogin":     "true",
		"/account/v1/tenantRegister": "true",
		"/health":                    "true",
	}
}

func checkPublicRoutes(req *http.Request) bool {
	allowd := GetPublicRoutes()
	public := false

	if _, ok := allowd[req.RequestURI]; ok {
		public = true
	}

	if !public {
		public = strings.HasPrefix(req.RequestURI, "/account/swagger/")
	}

	if !public {
		public = strings.HasPrefix(req.RequestURI, "/product/swagger/")
	}

	return public
}
