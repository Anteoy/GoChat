package websocket

import (
	webs "golang.org/x/net/websocket"
)

var Users = make([]*webs.Conn, 13, 20)

func Chat(ws *webs.Conn) {
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
