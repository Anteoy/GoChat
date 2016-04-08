package controller

import (
	"io"
	"log"
	ws "mynet/websocket"
	"net/http"
	"strconv"
)

type Msg struct {
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

func Chat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msgstring := r.Form["msg"][0]
	myid := GetMyId(r)
	msg := &Msg{Id: myid, Msg: msgstring}
	ws.SendAll(msg)
	io.WriteString(w, "success")

}

func ChatFromId(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msg := r.Form["msg"][0]
	id, err := strconv.Atoi(r.Form["fromid"][0])
	if err != nil {
		io.WriteString(w, "error")
		log.Println(err)
		return
	}
	err = ws.SendByUserId(id, msg)
	if err != nil {
		io.WriteString(w, "error")
		log.Println(err)
		return
	}
	io.WriteString(w, "success")
}
