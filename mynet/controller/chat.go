package controller

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
	ws "mynet/websocket"
	"net/http"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.Form["msg"][0]

	for _, webs := range ws.Users {
		err := websocket.JSON.Send(webs, msg)
		if err != nil {
			log.Panic(err)
		}
	}

	io.WriteString(w, "success")

}
