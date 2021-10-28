package service

import (
	"github.com/s51ds/n1mmweb/udp"
	"log"
)

func Statistic() {
	log.Println("Statistic service started")

	listenerChan := make(chan udp.QsoInfo)
	udp.ContactListener <- listenerChan
	for {
		select {
		case event := <-listenerChan:
			{
				log.Println(event)
			}
		}
	}
}
