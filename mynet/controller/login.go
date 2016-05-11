package controller

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mynet/redis"
	"net/http"
	"strconv"
	"strings"
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

	sessionId := GetSessionID(r)
	siCookie := &http.Cookie{Name: "GoSessionId", Value: sessionId, MaxAge: 0}

	cli := redis.RedisPool.Get()
	_, err := cli.Do("HSET", "sessionid", sessionId, id)

	if err == nil {
		http.SetCookie(w, siCookie)
	}

	return err
}

func GetSessionID(r *http.Request) string {
	ipAddr := r.RemoteAddr
	now := time.Now()
	nowTime := strconv.Itoa(now.Day()) + strconv.Itoa(now.Hour()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Second())
	sessionId := strconv.Itoa(createIpNum(ipAddr)) + nowTime

	h := md5.New()
	h.Write([]byte(sessionId))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)

}

func createIpNum(ipAddr string) int {
	ips := strings.Split(ipAddr, ".")
	if len(ips) == 1 {
		return 0
	}
	ipcount := 0
	for _, ip := range ips {
		ipToNum, err := strconv.Atoi(ip)
		if err != nil {
			ipToNum = 0
		}
		ipcount = ipcount + ipToNum
	}
	return ipcount
}
