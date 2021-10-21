package service

import (
	"fmt"
	"github.com/s51ds/n1mmweb/udp"
	locApi "github.com/s51ds/qthdb/api"
	distApi "github.com/s51ds/qthdist/api"
	"strings"
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
				resp := locApi.Locators(info.Call)
				sb := strings.Builder{}
				sb.WriteString(fmt.Sprintln(info.Call))
				if len(resp) > 0 {
					for _, v := range resp {
						if dist, azim, err := distApi.Distance("JN76TO", v.Locator); err != nil {
							fmt.Println(err)
						} else {
							sb.WriteString(fmt.Sprintf("%s %3dkm %3d° %s\n", v.Locator, int(dist), int(azim), v.LogTime.Sprint(false)))
						}
					}
				} else {
					sb.WriteString("NO LOCATORS IN DB")
				}

				fmt.Println(sb.String())

				//				qthdb.api.Locators(myLocator, info.Call)
			}
		}
	}

}
