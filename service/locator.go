package service

import (
	"fmt"
	"github.com/s51ds/n1mmweb/udp"
	"github.com/s51ds/n1mmweb/web"
	"github.com/s51ds/qthdb/locators"
	"github.com/s51ds/qthgeo/distance"
	"log"
	"strings"
)

func Locators(myLocator string) {
	log.Println("Locators service started")

	listenerChan := make(chan udp.QsoInfo)
	udp.LookupinfoListener <- listenerChan
	for {
		select {
		case info := <-listenerChan:
			{
				resp := locators.Get(info.Call)
				sb := strings.Builder{}
				sb.WriteString(fmt.Sprintln(info.Call) + "<br>")
				if len(resp) > 0 {
					for _, v := range resp {
						if dist, azim, err := distance.Get(myLocator, v.Locator); err != nil {
							fmt.Println(err)
						} else {
							sb.WriteString(fmt.Sprintf("%s %3dkm %3d° %s\n<br>", v.Locator, int(dist), int(azim), v.LogTime.Sprint(false)))
						}
					}
				} else {
					sb.WriteString("NO LOCATORS IN DB")
				}
				web.LocatorChan <- sb.String()
			}
		}
	}
}
