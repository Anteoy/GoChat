package controller

import (
	"io"
	"mynet/redis"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	paras := r.URL.Query()
	ids := paras["id"]
	if ids == nil {
		io.WriteString(w, "请输入账号")
		return
	}
	id := ids[0]
	passwds := paras["passwd"]
	if passwds == nil {
		io.WriteString(w, "请输入密码")
		return
	}
	passwd := passwds[0]
	passwdForRedis := redis.Get("user:" + id + ":pass")

	if passwd == passwdForRedis {
		http.ServeFile(w, r, "./static/html/index.html")
	} else {
		http.ServeFile(w, r, "./static/html/login.html")
	}
}
