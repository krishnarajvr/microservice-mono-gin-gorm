package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func init() {
	fmt.Println("krakend-example handler plugin loaded!!!")
}

func main() {}

// HandlerRegisterer is the name of the symbol krakend looks up to try and register plugins
var HandlerRegisterer registrable = registrable("krakend-example")

type registrable string

const outputHeaderName = "X-AUTH-TOKEN"
const pluginName = "krakend-example"

func (r registrable) RegisterHandlers(f func(
	name string,
	handler func(
		context.Context,
		map[string]interface{},
		http.Handler) (http.Handler, error),
)) {
	f(pluginName, r.registerHandlers)
}

func (r registrable) registerHandlers(ctx context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {
	//fmt.Println("PluginConf: ", pluginConf)

	fmt.Println("krakend-example is registered!")

	//client := &http.Client{Timeout: 3 * time.Second}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Validate Basic Auth
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		r2 := new(http.Request)
		*r2 = *r

		fmt.Println(r)
		fmt.Println(w)

		if auth[0] == "Basic" {
			fmt.Println("Handler calling with auth")
			fmt.Fprintf(w, "{ \"message\": \"")
			r2.Header.Set("X-CUSTOM-DATA", "My Custom Data with Auth")
			w.Header().Set("X-CUSTOM-DATA", "Writer My Custom Data with Auth")
			handler.ServeHTTP(w, r2)
			fmt.Fprintf(w, "\"}")

		} else {
			// Do nothing
			fmt.Println("Handler calling without Auth!")
			r.Header.Set("X-CUSTOM-DATA", "My Custom Data without Auth")
			w.Header().Set("X-CUSTOM-DATA", "Writer My Custom Data without Auth")
			handler.ServeHTTP(w, r)
		}

	}), nil
}
