package statistic

import "github.com/s51ds/n1mmweb/udp"

func isDupe(qso udp.QsoInfo) bool {
	return qso.Points == "0"
}
