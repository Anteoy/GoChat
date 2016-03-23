package controller

import (
	"fmt"
	"io"
	"mynet/redis"
	"net/http"
)

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, req.RequestURI)
}

func MyUri(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ok")
}

func Redis(res http.ResponseWriter, req *http.Request) {
}

func Hello(w http.ResponseWriter, r *http.Request) {

	paras := r.URL.Query()
	mingl := paras["m"][0]

	client := redis.RedisPool.Get()
	defer client.Close()

	v, _ := client.Do("GET", mingl)
	fmt.Printf("hello %s\n", v)
	fmt.Fprintf(w, "Hello! %s", v)
}
