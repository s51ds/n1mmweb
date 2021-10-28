package web

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Start(addr string) {
	fmt.Println("Web server started on ", addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
