package websocket

import (
	"fmt"
	"mynet/redis"

	gredis "github.com/garyburd/redigo/redis"
	webs "golang.org/x/net/websocket"
)

var Users = make([]*UserWebsocket, 0, 20)

type UserWebsocket struct {
	UserId int
	Ws     *webs.Conn
}

func DeleteUserFromUsers(user *UserWebsocket) {
	for i, u := range Users {
		if u == user {
			index := i + 1
			user.Ws.Close()
			Users = append(Users[:i], Users[index:]...)
			return
		}
	}
}

func Chat(ws *webs.Conn) {

	r := ws.Request()
	thisCookie, _ := r.Cookie("GoSessionId")
	cookie := thisCookie.Value
	userid, _ := gredis.Int(redis.Do("HGET", "sessionid", cookie))
	user := &UserWebsocket{UserId: userid, Ws: ws}

	fmt.Println(Users)
	for {
		for _, User := range Users {
			if ws == User.Ws && userid == User.UserId {
				goto forend
			}
			if userid == User.UserId && ws != User.Ws {
				DeleteUserFromUsers(User)
			}
		}
		Users = append(Users, user)
	forend:
	}
}
