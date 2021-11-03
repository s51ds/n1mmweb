package udp

import (
	"fmt"
	"log"
)

var (
	lookupInfoListeners = make(map[chan QsoInfo]bool)
	contactListeners    = make(map[chan QsoInfo]bool)
)

func Broadcaster() {
	fmt.Println("udp Broadcaster started")
	//lookupInfoListeners := make(map[chan QsoInfo]bool)
	//contactListeners := make(map[chan QsoInfo]bool)
	for {
		select {
		case info := <-LookupInfoChan:
			{
				for k, v := range lookupInfoListeners {
					if v {
						k <- info
					}
				}
			}
		case event := <-ContactInfoChan:
			{
				//				log.Println("ContactInfoChan", event.String())
				broadcastContactEvent(event)
			}
		case event := <-ContactReplaceChan:
			{
				//				log.Println("ContactReplaceChan", event.String())
				broadcastContactEvent(event)
			}
		case event := <-ContactDeleteChan:
			{
				//				log.Println("ContactDeleteChan", event.String())
				broadcastContactEvent(event)
			}
		case lookupListener := <-LookupinfoListener:
			{
				lookupInfoListeners[lookupListener] = true
				log.Println("LookupinfoListener registered, num of listener:", len(lookupInfoListeners))
			}
		case contactListener := <-ContactListener:
			{
				contactListeners[contactListener] = true
				log.Println("ContactListener registered, num of listeners:", len(contactListeners))
			}
		}

	}
}

func broadcastContactEvent(event QsoInfo) {
	for k, v := range contactListeners {
		if v {
			k <- event
		}
	}
}
