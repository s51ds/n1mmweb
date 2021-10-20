package udp

import "fmt"

func Broadcaster() {
	fmt.Println("Broadcaster started")
	lookupInfoListeners := make(map[chan QsoInfo]bool)
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
		case info := <-ContactInfoChan:
			{
				fmt.Println("ContactInfoChan", info.String())
			}
		case info := <-ContactReplaceChan:
			{
				fmt.Println("ContactReplaceChan", info.String())
			}
		case info := <-ContactDeleteChan:
			{
				fmt.Println("ContactDeleteChan", info.String())
			}
		case lookupListener := <-LookupinfoListener:
			{
				lookupInfoListeners[lookupListener] = true
				fmt.Println("LookupinfoListener registered, num of listener:", len(lookupInfoListeners))
			}
		}
	}
}
