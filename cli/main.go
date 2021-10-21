package main

import (
	"context"
	"fmt"
	"github.com/s51ds/n1mmweb/service"
	"github.com/s51ds/n1mmweb/udp"
	"github.com/s51ds/qthdb/db"
	"os"
)

func main() {

	fmt.Println("N1MM WEB started")
	if err := db.Open("db.gob"); err != nil {
		fmt.Println(err.Error())
		dir, _ := os.Getwd()
		fmt.Println("Working directory:", dir)

		os.Exit(1)
	}
	go udp.Broadcaster()
	go service.Locators("S59ABC")
	if err := udp.StartServer(context.Background(), "localhost:12060"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
