package web

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

func Start(addr string) {
	fmt.Println("Web server started on ", addr)
	setupRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	if _, err := fmt.Fprintf(w, index); err != nil {
		fmt.Println(err.Error())
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Tukaj se bodo prikazali lokatorji, ko bo čas za to"))

	webSocketWriter(ws)
}

var LocatorChan = make(chan (string))

func webSocketWriter(conn *websocket.Conn) {
	for {
		select {
		case locators := <-LocatorChan:
			{
				fmt.Println("webSocketWriter", locators)
				if err := conn.WriteMessage(1, []byte(locators)); err != nil {
					log.Println(err)
				}
			}
		}

		//if err := conn.WriteMessage(1, []byte(time.Now().String())); err != nil {
		//	log.Println(err)
		//	return
		//}
		//time.Sleep(time.Second)

	}
}

const index = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Locators</title>
</head>
<body>
<h2>LOCATORS</h2>

<h2><code id="locators"></code></h2>

<script>
    let socket = new WebSocket("ws://127.0.0.1:8080/ws");
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    socket.onmessage = event => {
        document.getElementById('locators').innerHTML = event.data
        console.log("Message received: ", event);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        socket.send("Client Closed!")
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
</script>
</body>
</html>
`
