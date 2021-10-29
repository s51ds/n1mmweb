package service

import (
	"github.com/s51ds/n1mmweb/udp"
	"log"
	"strings"
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
						// new Contact
						log.Println(event.ID, "contactInfo")
						qsoLogMu.Lock()
						qsoLog[event.ID] = event
						qsoLogMu.Unlock()
						log.Println("qsoLog", len(qsoLog))
						saveQsoLog()
					}
				case 4:
					{
						ID := strings.ReplaceAll(event.ID, "-", "")
						log.Println(ID, "contactDelete")
						qsoLogMu.Lock()
						delete(qsoLog, ID)
						qsoLogMu.Unlock()
						log.Println("qsoLog", len(qsoLog))
						saveQsoLog()
					}
				case 3:
					{
						log.Println(event.ID, "contactReplace")
						qsoLogMu.Lock()
						qsoLog[event.ID] = event
						qsoLogMu.Unlock()
						log.Println("qsoLog", len(qsoLog))
						saveQsoLog()
					}
				}
			}
		}
	}
}
