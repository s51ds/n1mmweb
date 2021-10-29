package statistic

import (
	"fmt"
	"github.com/s51ds/ctydb/api"
	"github.com/s51ds/n1mmweb/service"
	"log"
	"strconv"
	"strings"
)

// TODO: remove init
func init() {
	api.Load()
}

func Dxcc() string {
	data := make(map[string]int)

	service.QsoLogMux.RLock()
	for _, qso := range service.QsoLog {
		if points, err := strconv.Atoi(qso.Points); err != nil {
			log.Println("Qrb() dupe check", err)
		} else {
			if points == 0 { // dupe
				continue
			}
		} // no dupes

		if ctyDat, err := api.Get(qso.Call); err != nil {
			log.Println("Dxcc()", err.Error())
		} else {
			data[ctyDat.CountryName]++
		}
	}
	service.QsoLogMux.RUnlock()
	sb := strings.Builder{}
	for k, v := range data {
		sb.WriteString(fmt.Sprintf("%4d: %s\n", v, k))
	}
	return sb.String()
}
