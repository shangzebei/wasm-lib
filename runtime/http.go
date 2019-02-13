package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"wasmgo/types"
	"wasmgo/wasm"
)

type Http struct {
	types.VMInterface
}

//char * http_get(char * url);
func (h *Http) Http_get(url string) int64 {
	res, err := http.Get(url)
	if err != nil {
		return 0
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0
	}
	p := wasm.FindfreeSpece(h.Vm, int64(len(b)))
	copy(h.Vm.Memory[p:], b)
	return p
}

//char * http_get(char * url,char *contentType,char * body);
func (h *Http) Http_post(url string, contentType string, body string) int64 {
	fmt.Println(url, contentType, body)
	res, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		println(err)
		return 0
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println(err)
		return 0
	}
	p := wasm.FindfreeSpece(h.Vm, int64(len(b)))
	copy(h.Vm.Memory[p:], b)
	return p
}
