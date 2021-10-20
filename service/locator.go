package service

import (
	"fmt"
	"github.com/s51ds/n1mmweb/udp"
)

func Locators(myLocator string) {
	fmt.Println("Locators service started")

	listenerChan := make(chan udp.QsoInfo)
	udp.LookupinfoListener <- listenerChan
	for {
		select {
		case info := <-listenerChan:
			{
				fmt.Println("----->", info.Call)
			}
		}
	}

}
