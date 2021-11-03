package web

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/s51ds/n1mmweb/comm"
	"github.com/s51ds/n1mmweb/statistic"
	"log"
	"net/http"
	"os"
)

func Start(addr string) {
	fmt.Println("Web server started on ", addr)
	go webSocketWriterService()
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
	http.HandleFunc("/qrb", qrbPage)
	http.HandleFunc("/dxcc", dxccPage)
	http.HandleFunc("/dist", distPage)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	if _, err := fmt.Fprintf(w, index); err != nil {
		log.Println("homePage", err.Error())
	}
}

func qrbPage(w http.ResponseWriter, r *http.Request) {
	log.Println("qrbPage()", r.Host)
	//TODO cli context
	s := statistic.Qrb("JN76TO")
	if _, err := fmt.Fprintf(w, s); err != nil {
		log.Println("qrbPage()", err.Error())
	}
}

func dxccPage(w http.ResponseWriter, r *http.Request) {
	log.Println("dxccPage()", r.Host)
	s := statistic.Dxcc()
	if _, err := fmt.Fprintf(w, s); err != nil {
		log.Println("dxccPage()", err.Error())
	}
}

func distPage(w http.ResponseWriter, r *http.Request) {
	log.Println("distPage()", r.Host)
	//TODO cli context
	s := statistic.Dist("JN76TO")
	if _, err := fmt.Fprintf(w, s); err != nil {
		log.Println("distPage()", err.Error())
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("wsEndpoint.Upgrade", err)
	}

	log.Println(r.Host, "Client Connected", wsConnNum)
	if err = ws.WriteMessage(1, []byte("Tukaj se bodo prikazali lokatorji, ko bo čas za to")); err != nil {
		fmt.Println("wsEndpoint.WriteMessage", err.Error())
	} else {

		wsMap[wsConnNum] = ws
		wsConnNum++
	}

}

var (
	wsConnNum = 1
	wsMap     = make(map[int]*websocket.Conn)
	//LocatorChanWebServer = make(chan string)
)

func webSocketWriterService() {
	for {
		select {
		case locators := <-comm.LocatorChanWebServer:
			{
				for i, ws := range wsMap {
					if err := ws.WriteMessage(1, []byte(locators)); err != nil {
						fmt.Println("webSocketWriterService", i, err.Error())
						delete(wsMap, i)
					}
				}
			}
		}
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
