package websocket

import (
	"fmt"

	webs "golang.org/x/net/websocket"
)

var Users = make([]*webs.Conn, 0, 20)

func Chat(ws *webs.Conn) {
	fmt.Println(Users)
	for {
		for _, thisws := range Users {
			if ws == thisws {
				goto forend
			}
		}
		Users = append(Users, ws)
	forend:
	}
}
