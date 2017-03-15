package usewebsocket

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
)

func echo(ws *websocket.Conn) {
	var err error
	for {
		var replay string
		if err = websocket.Message.Receive(ws, &replay); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client:" + replay)
		msg := "Received: " + replay
		fmt.Println("Sending to client: " + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't Send")
			break
		}
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/usewebsocket/templates/usws.gtpl")
		t.Execute(w, nil)
	}
}

func UseEcho() {
	http.HandleFunc("/", login)
	http.Handle("/ws", websocket.Handler(echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
