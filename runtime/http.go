package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Http struct {
}

//char * http_get(char * url);
func (h *Http) Http_get(url string) string {
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
func (h *Http) Http_post(url string, contentType string, body string) string {
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
