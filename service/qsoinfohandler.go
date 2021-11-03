package service

import (
	"github.com/s51ds/n1mmweb/udp"
	"log"
	"strings"
)

func QsoInfoHandler() {
	log.Println("QsoInfoHandler service started")

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
						// new Contact
						//						log.Println(event.ID, "contactInfo")
						QsoLogMux.Lock()
						QsoLog[event.ID] = event
						QsoLogMux.Unlock()
						//						log.Println("QsoLog", len(QsoLog))
						saveQsoLog()
					}
				case 4:
					{
						ID := strings.ReplaceAll(event.ID, "-", "")
						//						log.Println(ID, "contactDelete")
						QsoLogMux.Lock()
						delete(QsoLog, ID)
						QsoLogMux.Unlock()
						//						log.Println("QsoLog", len(QsoLog))
						saveQsoLog()
					}
				case 3:
					{
						//						log.Println(event.ID, "contactReplace")
						QsoLogMux.Lock()
						QsoLog[event.ID] = event
						QsoLogMux.Unlock()
						//						log.Println("QsoLog", len(QsoLog))
						saveQsoLog()
					}
				}
			}
		}
	}
}
