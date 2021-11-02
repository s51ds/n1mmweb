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
		km000: "    < 100 km : ",
		km100: "100 - 200 km : ",
		km200: "200 - 300 km : ",
		km300: "300 - 400 km : ",
		km400: "400 - 500 km : ",
		km500: "500 - 600 km : ",
		km600: "600 - 700 km : ",
		km700: "700 - 800 km : ",
		km800: "    > 800 km : ",
	}
)

func Qrb(myLocator string) string {
	data := make([]int, len(qrbDesc), len(qrbDesc))

	// points
	sumPoints := 0
	points := make([]int, len(qrbDesc), len(qrbDesc))

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
			pts := int(dist)
			sumPoints += pts
			switch {
			case int(dist) < 100:
				{
					data[km000]++
					points[km000] += pts
				}
			case dist >= 100 && dist < 200:
				{
					data[km100]++
					points[km100] += pts
				}
			case dist >= 200 && dist < 300:
				{
					data[km200]++
					points[km200] += pts
				}
			case dist >= 300 && dist < 400:
				{
					data[km300]++
					points[km300] += pts
				}
			case dist >= 400 && dist < 500:
				{
					data[km400]++
					points[km400] += pts
				}
			case dist >= 500 && dist < 600:
				{
					data[km500]++
					points[km500] += pts
				}
			case dist >= 600 && dist < 700:
				{
					data[km600]++
					points[km600] += pts
				}
			case dist >= 700 && dist < 800:
				{
					data[km700]++
					points[km700] += pts
				}
			case dist >= 800:
				{
					data[km800]++
					points[km800] += pts
				}
			}
		}
	} // for
	service.QsoLogMux.RUnlock()

	sb := strings.Builder{}
	for i, v := range data {
		ptsPct := float64(points[i]) / float64(sumPoints) * 100
		sb.WriteString(fmt.Sprintf("%s%3d QSO; %6d pts -> %4.1f", qrbDesc[i], v, points[i], ptsPct))
		sb.WriteString(" %%\n")
	}
	return sb.String()
}
