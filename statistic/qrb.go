package statistic

import (
	"fmt"
	"github.com/s51ds/n1mmweb/service"
	"github.com/s51ds/qthgeo/distance"
	"log"
	"strconv"
	"strings"
)

const (
	km000 int = iota
	km100
	km200
	km300
	km400
	km500
	km600
	km700
	km800
)

var (
	qrbDesc = map[int]string{
		km000: "< 100 km",
		km100: "100 - 200 km",
		km200: "200 - 300 km",
		km300: "300 - 400 km",
		km400: "400 - 500 km",
		km500: "500 - 600 km",
		km600: "600 - 700 km",
		km700: "700 - 800 km",
		km800: "> 800 km",
	}
)

func Qrb() string {
	data := make([]int, len(qrbDesc), len(qrbDesc))

	//TODO cli context
	myLocator := "JN76TO"

	service.QsoLogMux.RLock()
	for _, qso := range service.QsoLog {
		if points, err := strconv.Atoi(qso.Points); err != nil {
			log.Println("Qrb() dupe check", err)
		} else {
			if points == 0 { // dupe
				continue
			}
		}

		if dist, _, err := distance.Get(myLocator, qso.GridSquare); err != nil {
			log.Println("Qrb()", err.Error())
		} else {
			switch {
			case dist < 100:
				{
					data[km000]++
				}
			case dist >= 100 && dist < 200:
				{
					data[km100]++
				}
			case dist >= 200 && dist < 300:
				{
					data[km200]++
				}
			case dist >= 300 && dist < 400:
				{
					data[km300]++
				}
			case dist >= 400 && dist < 500:
				{
					data[km400]++
				}
			case dist >= 500 && dist < 600:
				{
					data[km500]++
				}
			case dist >= 600 && dist < 700:
				{
					data[km600]++
				}
			case dist >= 700 && dist < 800:
				{
					data[km700]++
				}
			case dist >= 800:
				{
					data[km800]++
				}
			}
		}
	} // for
	service.QsoLogMux.RUnlock()

	sb := strings.Builder{}
	for i, v := range data {
		sb.WriteString(fmt.Sprintf("%4d: %s\n", v, qrbDesc[i]))
	}
	return sb.String()
}
