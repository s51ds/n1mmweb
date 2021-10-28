package udp

var (
	LookupInfoChan     = make(chan QsoInfo)
	ContactInfoChan    = make(chan QsoInfo)
	ContactReplaceChan = make(chan QsoInfo)
	ContactDeleteChan  = make(chan QsoInfo)

	LookupinfoListener = make(chan chan QsoInfo)
)
