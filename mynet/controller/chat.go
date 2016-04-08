package controller

import (
	"io"
	"log"
	ws "mynet/websocket"
	"net/http"

	"golang.org/x/net/websocket"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.Form["msg"][0]

	for _, user := range ws.Users {
		if user == nil {
			continue
		}
		err := websocket.JSON.Send(user.Ws, msg)
		if err != nil {
			log.Panic(err)
			io.WriteString(w, "websocket is err!")
			return
		}
	}

	io.WriteString(w, "success")

}
