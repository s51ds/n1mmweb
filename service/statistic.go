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
				switch event.Type {
				case 2:
					{
						log.Println(event.ID, "contactInfo")
					}
				case 4:
					{
						log.Println(event.ID, "contactDelete")

					}
				case 3:
					{
						log.Println(event.ID, "contactReplace")

					}

				}
			}
		}
	}
}
