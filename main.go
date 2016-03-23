package main

import (
	"fmt"
	"mynet"
	"net/http"
	"os"
)

func main() {

	mynet.InitPool()

	server := http.NewServeMux()
	staticDirHandler(server, "/static", "./static", 0)
	server.HandleFunc("/redis", mynet.Hello)
	server.HandleFunc("/hello", mynet.HelloHandler)
	server.HandleFunc("/niaho", mynet.MyUri)
	server.HandleFunc("/", index)
	err := http.ListenAndServe(":10000", server)
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
