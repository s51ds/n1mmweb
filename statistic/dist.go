package statistic

import (
	"fmt"
	"github.com/s51ds/n1mmweb/service"
	"github.com/s51ds/qthgeo/distance"
	"log"
	"sort"
	"strconv"
	"strings"
)

type qsoDistItem struct {
	dist     int
	azim     string
	locator  string
	callSign string
}

func (q qsoDistItem) String() string {
	return fmt.Sprintf("%4d km %3s° %6s %s", q.dist, q.azim, q.locator, q.callSign)
}

type byDist []qsoDistItem

func (b byDist) Len() int {
	return len(b)
}

func (b byDist) Less(i, j int) bool {
	return b[i].dist > b[j].dist
}

func (b byDist) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func Dist(myLocator string) string {

	service.QsoLogMux.RLock()
	data := make(byDist, 0, len(service.QsoLog))
	for _, qso := range service.QsoLog {
		if isDupe(qso) {
			continue
		}

		if dist, azim, err := distance.Get(myLocator, qso.GridSquare); err != nil {
			log.Println("Qrb()", err.Error())
		} else {
			d := qsoDistItem{
				dist:     int(dist),
				azim:     strconv.Itoa(int(azim)),
				locator:  qso.GridSquare,
				callSign: qso.Call,
			}
			data = append(data, d)
		}
	}
	service.QsoLogMux.RUnlock()

	sort.Sort(byDist(data))

	sb := strings.Builder{}
	for _, k := range data {
		sb.WriteString(k.String() + "\n")
	}
	return sb.String()
}
