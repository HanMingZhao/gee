package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

//Engine类型实现http.Handler接口的ServerHTTP方法
func (engine *Engine) ServeHTTP(wri http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(wri, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(wri, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(wri, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
