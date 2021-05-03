package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
var ClientRegisterer = registerer("proxy-plugin")

type registerer string

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), r.registerClients)
}

func (r registerer) registerClients(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
	// check the passed configuration and initialize the plugin
	name, ok := extra["name"].(string)

	if !ok {
		return nil, errors.New("wrong config")
	}

	if name != string(r) {
		return nil, fmt.Errorf("unknown register %s", name)
	}

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http client
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("proxy-plugin called")
		fmt.Println("Request: ", req)
		fmt.Println("Request Header: ", req.Header)
		fmt.Println("Tenant Id: ", req.Header.Get("tenantId"))
		fmt.Println("User Agent: ", req.UserAgent())
		fmt.Println("Reqeust Context: ", req.Context())
		fmt.Println("Context: ", context.Background)
		fmt.Println("Extra: ", extra)

		client := &http.Client{
			Timeout: time.Second * 10,
		}

		newResp, err := client.Do(req)

		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		defer newResp.Body.Close()

		body, err := ioutil.ReadAll(newResp.Body)

		fmt.Println("End calling proxy plugin")

		w.Write(body)
	}), nil
}

func init() {
	fmt.Println("proxy-plugin client plugin loaded!!!")
}

func main() {}
