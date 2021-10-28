package main

import (
	"context"
	"fmt"
	"github.com/s51ds/n1mmweb/service"
	"github.com/s51ds/n1mmweb/udp"
	"github.com/s51ds/n1mmweb/web"
	"github.com/s51ds/qthdb/db"
	"log"
	"os"
)

func main() {
	log.Println("N1MM WEB started")

	myLocator := "JN76TO"
	udpSocket := "localhost:12060"
	webSocket := ":8080"

	go web.Start(webSocket)

	if err := db.Open("db.gob"); err != nil {
		fmt.Println(err.Error())
		dir, _ := os.Getwd()
		fmt.Println("Working directory:", dir)

		os.Exit(1)
	}
	fmt.Println("My Locator", myLocator)
	fmt.Println("listen on udp socket", udpSocket)

	go udp.Broadcaster()
	go service.Locators(myLocator)
	go service.Statistic()

	if err := udp.StartServer(context.Background(), udpSocket); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

}
