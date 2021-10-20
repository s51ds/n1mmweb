package main

import (
	"context"
	"fmt"
	"github.com/s51ds/n1mmweb/udp"
	"os"
)

func main() {

	fmt.Println("N1MM WEB started")
	if err := udp.StartServer(context.Background(), "localhost:12060"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
