package lib

import (
	"wasmgo/types"
	"wasmgo/wasm"
)

type Http struct {
	types.RegInterface
}

const NETWORK = "NETWORK"

//char * http_get(char * url);
func (h *Http) Http_get(url string) string {
	re := wasm.PlugInstants(NETWORK).Call("Http_get", url)
	if re != nil {
		return re[0].(string)
	}
	return "null"
}

//char * http_get(char * url,char *contentType,char * body);
func (h *Http) Http_post(url string, contentType string, body string) string {
	re := wasm.PlugInstants(NETWORK).Call("Http_post", url, contentType, body)
	if re != nil {
		return re[0].(string)
	}
	return "null"
}
