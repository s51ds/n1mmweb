package service

import (
	"fmt"
	"github.com/s51ds/n1mmweb/udp"
	"github.com/s51ds/qthdb/api"
)

var data string

func Locators(myLocator string) {
	fmt.Println("Locators service started")

	listenerChan := make(chan udp.QsoInfo)
	udp.LookupinfoListener <- listenerChan
	for {
		select {
		case info := <-listenerChan:
			{
				fmt.Println("----->", info.Call)
				resp := api.Locators(info.Call)
				fmt.Println(resp)

				//				qthdb.api.Locators(myLocator, info.Call)
			}
		}
	}

}
