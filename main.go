package main

import (
	"fmt"
	c "mynet/controller"
	_ "mynet/redis"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/", index)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(":2224", nil)
	if err != nil {
		fmt.Println("bind error")
	}
}

func staticDirHandler(server *http.ServeMux, prefix string, staticDir string, flags int) {
	server.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if e := isExists(file); !e {
			http.NotFound(w, r)
			fmt.Println("未找到" + file)
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
	http.ServeFile(w, r, "./static/html/login.html")
}
