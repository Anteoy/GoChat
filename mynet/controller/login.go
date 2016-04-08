package controller

import (
	"io"
	"mynet/redis"
	"net/http"
	"strconv"
	"time"

	gredis "github.com/garyburd/redigo/redis"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ids := r.Form["id"]
	if ids == nil {
		io.WriteString(w, "请输入账号")
		return
	}
	id := ids[0]
	passwds := r.Form["passwd"]
	if passwds == nil {
		io.WriteString(w, "请输入密码")
		return
	}
	passwd := passwds[0]
	passwdForRedis := redis.Get("user:" + id + ":pass")

	if passwd == passwdForRedis {
		if err := addSessionId(w, r, id); err != nil {
			goto addSIIsError
		}
		http.ServeFile(w, r, "./static/html/index.html")
	} else {
		http.ServeFile(w, r, "./static/html/login.html")
	}
addSIIsError:
	http.ServeFile(w, r, "./static/html/login.html")
}

func MyId(w http.ResponseWriter, r *http.Request) {
	userid := GetMyId(r)
	io.WriteString(w, strconv.Itoa(userid))
}

func GetMyId(r *http.Request) int {
	thisCookie, _ := r.Cookie("GoSessionId")
	cookie := thisCookie.Value
	userid, _ := gredis.Int(redis.Do("HGET", "sessionid", cookie))
	return userid
}

func addSessionId(w http.ResponseWriter, r *http.Request, id string) error {
	ip := r.RemoteAddr
	now := time.Now()
	sessionId := ip + strconv.Itoa(now.Day()) + strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	siCookie := &http.Cookie{Name: "GoSessionId", Value: sessionId, MaxAge: 0}

	cli := redis.RedisPool.Get()
	_, err := cli.Do("HSET", "sessionid", sessionId, id)

	if err == nil {
		http.SetCookie(w, siCookie)
	}

	return err
}
