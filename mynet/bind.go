package mynet

import (
	"io"
	"net/http"
)

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, req.RequestURI)
}

func MyUri(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ok")
}
