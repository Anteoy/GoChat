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

func Redis(w http.ResponseWriter, r *http.Request) {
	paras := r.URL.Query()
	mingl := paras["m"][0]

	v := redis.Get(mingl)

	fmt.Printf("hello %s\n", v)
	fmt.Fprintf(w, "Hello! %s", v)
}

func Login(w http.ResponseWriter, r *http.Request) {

	paras := r.URL.Query()
	id := paras["id"][0]
	passwd := paras["passwd"][0]

	passwdForRedis := redis.Get("user:" + id + ":pass")

	if passwd == passwdForRedis {
		http.ServeFile(w, r, "./static/html/login.html")
	} else {
		http.ServeFile(w, r, "./static/html/index.html")
	}
}
