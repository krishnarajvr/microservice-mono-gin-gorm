package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// HandlerRegisterer is the symbol the plugin loader will try to load. It must implement the Registerer interface
var HandlerRegisterer = registerer("router-plugin")

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func validateToken(token string) bool {
	return true
}

//registerHandlers Executes before every requests
func (r registerer) registerHandlers(ctx context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {
	name, ok := extra["name"].([]interface{})

	if !ok {
		return nil, errors.New("wrong config")
	}

	if name[0] != string(r) {
		return nil, fmt.Errorf("unknown register %s", name)
	}

	// return the actual handler wrapping with custom logic
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		public := checkPublicRoutes(req)

		if public {
			handler.ServeHTTP(w, req)
			return
		}

		context, authErr := GetAuthorizeContext(req)

		if authErr != nil {
			fmt.Println("Auth error:", authErr)
			SendError(w, "Invalid User Token", http.StatusForbidden)
			return
		}

		if context.Status != 200 {
			SendError(w, context.Error.Message, http.StatusForbidden)
			return
		}

		//Set Headers for upstream services
		req.Header.Set("X-Tenant-Id", context.Data.Auth.TenantId)

		if len(context.Data.Auth.UserId) > 0 {
			req.Header.Set("X-User-Id", context.Data.Auth.UserId)
		}

		if len(context.Data.Auth.Type) > 0 {
			req.Header.Set("X-Auth-Type", context.Data.Auth.Type)
		}

		if len(context.Data.Auth.ReferenceId) > 0 {
			req.Header.Set("X-Reference-Id", context.Data.Auth.ReferenceId)
		}

		handler.ServeHTTP(w, req)

		return
	}), nil
}

func SendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(AccessDenied(message))
}

func init() {
	//Todo - debug purpose - need to remove
	fmt.Println("router-plugin: Init handler loaded!!!")
}

func main() {}
