package statistic

import (
	"fmt"
	"github.com/s51ds/ctydb/api"
	"github.com/s51ds/n1mmweb/service"
	"log"
	"sort"
	"strconv"
	"strings"
)

// TODO: remove init
func init() {
	api.Load()
}

type qsoDxccItem struct {
	prefix string
	qsos   int
	pts    int
}

func (a qsoDxccItem) String() string {
	return fmt.Sprintf("%3s %4d QSO; %6d pts", a.prefix, a.qsos, a.pts)
}

type byDxcc []qsoDxccItem

func (b byDxcc) Len() int {
	return len(b)
}

func (b byDxcc) Less(i, j int) bool {
	return b[i].prefix < b[j].prefix
}

func (b byDxcc) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func Dxcc() string {
	// points
	sumPoints := 0

	service.QsoLogMux.RLock()
	m := make(map[string]qsoDxccItem)

	for _, qso := range service.QsoLog {
		if isDupe(qso) {
			continue
		}
		ctyDat, err := api.Get(qso.Call)
		if err != nil {
			log.Println("Dxcc()", err.Error())
			continue
		}
		points, err := strconv.Atoi(qso.Points)
		if err != nil {
			log.Println("Dxcc()", err.Error())
			continue
		}
		sumPoints += points
		pfx := ctyDat.PrimaryPrefix
		v, has := m[pfx]
		if has {
			v.qsos++
			v.pts += points
			m[pfx] = v
		} else {
			v := qsoDxccItem{
				prefix: pfx,
				qsos:   1,
				pts:    points,
			}
			m[pfx] = v
		}

	}
	service.QsoLogMux.RUnlock()

	list := make([]qsoDxccItem, 0, len(m))
	for k, v := range m {
		v.prefix = k
		list = append(list, v)
	}

	sort.Sort(byDxcc(list))

	sb := strings.Builder{}
	for _, v := range list {
		pct := float64(v.pts) / float64(sumPoints) * 100
		sb.WriteString(fmt.Sprintf("%s -> %6.1f ", v.String(), pct))
		sb.WriteString("%%\n")
	}
	return sb.String()
}
