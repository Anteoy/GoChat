package main

import (
	"fmt"
	c "mynet/controller"
	"mynet/redis"
	"net/http"
	"os"
)

func main() {

	redis.InitPool()

	server := http.NewServeMux()
	staticDirHandler(server, "/static", "./static", 0)
	server.HandleFunc("/redis", c.Redis)
	server.HandleFunc("/hello", c.HelloHandler)
	server.HandleFunc("/nihao", c.MyUri)
	server.HandleFunc("/login", c.Login)
	server.HandleFunc("/", index)
	err := http.ListenAndServe(":1111", server)
	if err != nil {
		fmt.Println("bind error")
	}
}

func staticDirHandler(server *http.ServeMux, prefix string, staticDir string, flags int) {
	server.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if e := isExists(file); !e {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, file)
	})
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/index.html")
}
