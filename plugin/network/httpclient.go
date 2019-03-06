package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"wasmgo/types"
)

func Init(plugin *types.VMPlugin) {
	plugin.Reg("Http_get")
	plugin.Reg("Http_post")
	plugin.Reg("Hello")

	plugin.PlugName = "NETWORK"
	plugin.Version = "0.0.1"

}

func Http_get(url string) string {
	res, err := http.Get(url)
	if err != nil {
		return ""
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(b)
}

//char * http_get(char * url,char *contentType,char * body);
func Http_post(url string, contentType string, body string) string {
	fmt.Println(url, contentType, body)
	res, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		println(err)
		return ""
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println(err)
		return ""
	}
	return string(b)
}

func Hello() {
	fmt.Println("good")
}
