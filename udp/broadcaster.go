package udp

import "fmt"

func Broadcaster() {
	//	listeners := make(map[QsoInfoChan]bool)
	for {
		select {
		case QsoInfo := <-LookupInfoChan:
			{
				fmt.Println("LookupInfoChan", QsoInfo.String())
			}
		case QsoInfo := <-ContactInfoChan:
			{
				fmt.Println("ContactInfoChan", QsoInfo.String())
			}
		case QsoInfo := <-ContactReplaceChan:
			{
				fmt.Println("ContactReplaceChan", QsoInfo.String())
			}
		case QsoInfo := <-ContactDeleteChan:
			{
				fmt.Println("ContactDeleteChan", QsoInfo.String())
			}
		}
	}
}
