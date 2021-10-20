package udp

import (
	"encoding/json"
	"encoding/xml"
	"log"
)

const (
	lookupInfo     = "lookupinfo"
	contactInfo    = "contactinfo"
	contactReplace = "contactreplace"
	contactDelete  = "contactdelete"
)

type QsoInfoType int

const (
	unknown = iota
	lookupinfo
	contactinfo
	contactreplace
	contactdelete
)

var qsoInfoTypeNames = [...]string{"unknown", "lookupinfo", "contactinfo", "contactreplace", "contactdelete"}

func tokenNameToQsoInfoType(tokenName string) QsoInfoType {
	for i, v := range qsoInfoTypeNames {
		if v == tokenName {
			return QsoInfoType(i)
		}
	}
	log.Printf("Unexpected token: %s\n", tokenName)
	return QsoInfoType(0)
}

type QsoInfo struct {
	Type            QsoInfoType
	Timestamp       string `xml:"timestamp"`
	MyCall          string `xml:"mycall"`
	Band            string `xml:"band"`
	Mode            string `xml:"mode"`
	Call            string `xml:"call"`
	Sntnr           string `xml:"sntnr"`
	Rcvnr           string `xml:"rcvnr"`
	GridSquare      string `xml:"gridsquare"`
	Points          string `xml:"points"`
	RadioNr         string `xml:"radionr"`
	NetworkedCompNr string `xml:"NetworkedCompNr"`
	NetBiosName     string `xml:"NetBiosName"`
	StationName     string `xml:"StationName"`
	ID              string `xml:"ID"`
}

func (a *QsoInfo) String() string {
	if bb, err := json.Marshal(a); err != nil {
		return err.Error()
	} else {
		return qsoInfoTypeNames[a.Type] + ":" + string(bb)
	}
}

//func (a *QsoInfo) String() string {
//	return fmt.Sprintf("QsoInfo{Type:%s, Call:%s, GridSquare:%s}\n", qsoInfoTypeNames[a.Type], a.Call, a.GridSquare )
//}

type AppInfo struct {
	XMLName     xml.Name `xml:"AppInfo"`
	Text        string   `xml:",chardata"`
	App         string   `xml:"app"`
	Dbname      []string `xml:"dbname"`
	Contestnr   string   `xml:"contestnr"`
	Contestname string   `xml:"contestname"`
	StationName string   `xml:"StationName"`
}

type ContactInfo struct {
	//XMLName         xml.Name `xml:"contactinfo"`
	//Text            string   `xml:",chardata"`
	App             string `xml:"app"`
	Contestname     string `xml:"contestname"`
	Contestnr       string `xml:"contestnr"`
	Timestamp       string `xml:"timestamp"`
	Mycall          string `xml:"mycall"`
	Band            string `xml:"band"`
	Rxfreq          string `xml:"rxfreq"`
	Txfreq          string `xml:"txfreq"`
	Operator        string `xml:"operator"`
	Mode            string `xml:"mode"`
	Call            string `xml:"call"`
	Countryprefix   string `xml:"countryprefix"`
	Wpxprefix       string `xml:"wpxprefix"`
	Stationprefix   string `xml:"stationprefix"`
	Continent       string `xml:"continent"`
	Snt             string `xml:"snt"`
	Sntnr           string `xml:"sntnr"`
	Rcv             string `xml:"rcv"`
	Rcvnr           string `xml:"rcvnr"`
	Gridsquare      string `xml:"gridsquare"`
	Exchangel       string `xml:"exchangel"`
	Section         string `xml:"section"`
	Comment         string `xml:"comment"`
	Qth             string `xml:"qth"`
	Name            string `xml:"name"`
	Power           string `xml:"power"`
	Misctext        string `xml:"misctext"`
	Zone            string `xml:"zone"`
	Prec            string `xml:"prec"`
	Ck              string `xml:"ck"`
	Ismultiplierl   string `xml:"ismultiplierl"`
	Ismultiplier2   string `xml:"ismultiplier2"`
	Ismultiplier3   string `xml:"ismultiplier3"`
	Points          string `xml:"points"`
	Radionr         string `xml:"radionr"`
	RoverLocation   string `xml:"RoverLocation"`
	RadioInterfaced string `xml:"RadioInterfaced"`
	NetworkedCompNr string `xml:"NetworkedCompNr"`
	IsOriginal      string `xml:"IsOriginal"`
	NetBiosName     string `xml:"NetBiosName"`
	IsRunQSO        string `xml:"IsRunQSO"`
	StationName     string `xml:"StationName"`
	ID              string `xml:"ID"`
	IsClaimedQso    string `xml:"IsClaimedQso"`
}

type ContactReplace struct {
	XMLName         xml.Name `xml:"contactreplace"`
	Text            string   `xml:",chardata"`
	App             string   `xml:"app"`
	Contestname     string   `xml:"contestname"`
	Contestnr       string   `xml:"contestnr"`
	Timestamp       string   `xml:"timestamp"`
	Mycall          string   `xml:"mycall"`
	Band            string   `xml:"band"`
	Rxfreq          string   `xml:"rxfreq"`
	Txfreq          string   `xml:"txfreq"`
	Operator        string   `xml:"operator"`
	Mode            string   `xml:"mode"`
	Call            string   `xml:"call"`
	Countryprefix   string   `xml:"countryprefix"`
	Wpxprefix       string   `xml:"wpxprefix"`
	Stationprefix   string   `xml:"stationprefix"`
	Continent       string   `xml:"continent"`
	Snt             string   `xml:"snt"`
	Sntnr           string   `xml:"sntnr"`
	Rcv             string   `xml:"rcv"`
	Rcvnr           string   `xml:"rcvnr"`
	Gridsquare      string   `xml:"gridsquare"`
	Exchangel       string   `xml:"exchangel"`
	Section         string   `xml:"section"`
	Comment         string   `xml:"comment"`
	Qth             string   `xml:"qth"`
	Name            string   `xml:"name"`
	Power           string   `xml:"power"`
	Misctext        string   `xml:"misctext"`
	Zone            string   `xml:"zone"`
	Prec            string   `xml:"prec"`
	Ck              string   `xml:"ck"`
	Ismultiplierl   string   `xml:"ismultiplierl"`
	Ismultiplier2   string   `xml:"ismultiplier2"`
	Ismultiplier3   string   `xml:"ismultiplier3"`
	Points          string   `xml:"points"`
	Radionr         string   `xml:"radionr"`
	RoverLocation   string   `xml:"RoverLocation"`
	RadioInterfaced string   `xml:"RadioInterfaced"`
	NetworkedCompNr string   `xml:"NetworkedCompNr"`
	IsOriginal      string   `xml:"IsOriginal"`
	NetBiosName     string   `xml:"NetBiosName"`
	IsRunQSO        string   `xml:"IsRunQSO"`
	StationName     string   `xml:"StationName"`
	ID              string   `xml:"ID"`
	IsClaimedQso    string   `xml:"IsClaimedQso"`
}

type ContactDelete struct {
	XMLName     xml.Name `xml:"contactdelete"`
	Text        string   `xml:",chardata"`
	App         string   `xml:"app"`
	Timestamp   string   `xml:"timestamp"`
	Call        string   `xml:"call"`
	Contestnr   string   `xml:"contestnr"`
	StationName string   `xml:"StationName"`
	ID          string   `xml:"ID"`
}

type LookupInfo struct {
	//XMLName         xml.Name `xml:"lookupinfo"`
	//Text            string   `xml:",chardata"`
	App             string `xml:"app"`
	Contestname     string `xml:"contestname"`
	Contestnr       string `xml:"contestnr"`
	Timestamp       string `xml:"timestamp"`
	Mycall          string `xml:"mycall"`
	Band            string `xml:"band"`
	Rxfreq          string `xml:"rxfreq"`
	Txfreq          string `xml:"txfreq"`
	Operator        string `xml:"operator"`
	Mode            string `xml:"mode"`
	Call            string `xml:"call"`
	Countryprefix   string `xml:"countryprefix"`
	Wpxprefix       string `xml:"wpxprefix"`
	Stationprefix   string `xml:"stationprefix"`
	Continent       string `xml:"continent"`
	Snt             string `xml:"snt"`
	Sntnr           string `xml:"sntnr"`
	Rcv             string `xml:"rcv"`
	Rcvnr           string `xml:"rcvnr"`
	Gridsquare      string `xml:"gridsquare"`
	Exchangel       string `xml:"exchangel"`
	Section         string `xml:"section"`
	Comment         string `xml:"comment"`
	Qth             string `xml:"qth"`
	Name            string `xml:"name"`
	Power           string `xml:"power"`
	Misctext        string `xml:"misctext"`
	Zone            string `xml:"zone"`
	Prec            string `xml:"prec"`
	Ck              string `xml:"ck"`
	Ismultiplierl   string `xml:"ismultiplierl"`
	Ismultiplier2   string `xml:"ismultiplier2"`
	Ismultiplier3   string `xml:"ismultiplier3"`
	Points          string `xml:"points"`
	Radionr         string `xml:"radionr"`
	RoverLocation   string `xml:"RoverLocation"`
	RadioInterfaced string `xml:"RadioInterfaced"`
	NetworkedCompNr string `xml:"NetworkedCompNr"`
	IsOriginal      string `xml:"IsOriginal"`
	NetBiosName     string `xml:"NetBiosName"`
	IsRunQSO        string `xml:"IsRunQSO"`
	StationName     string `xml:"StationName"`
	ID              string `xml:"ID"`
	IsClaimedQso    string `xml:"IsClaimedQso"`
}

type RadioInfo struct {
	XMLName         xml.Name `xml:"RadioInfo"`
	Text            string   `xml:",chardata"`
	App             string   `xml:"app"`
	StationName     string   `xml:"StationName"`
	RadioNr         string   `xml:"RadioNr"`
	Freq            string   `xml:"Freq"`
	TXFreq          string   `xml:"TXFreq"`
	Mode            string   `xml:"Mode"`
	OpCall          string   `xml:"OpCall"`
	IsRunning       string   `xml:"IsRunning"`
	FocusEntry      string   `xml:"FocusEntry"`
	EntryWindowHwnd struct {
		Text               string `xml:",chardata"`
		Antenna            string `xml:"Antenna"`
		Rotors             string `xml:"Rotors"`
		FocusRadioNr       string `xml:"FocusRadioNr"`
		IsStereo           string `xml:"IsStereo"`
		IsSplit            string `xml:"IsSplit"`
		ActiveRadioNr      string `xml:"ActiveRadioNr"`
		IsTransmitting     string `xml:"IsTransmitting"`
		FunctionKeyCaption string `xml:"FunctionKeyCaption"`
		RadioName          string `xml:"RadioName"`
	} `xml:"EntryWindowHwnd"`
}

type Spot struct {
	XMLName     xml.Name `xml:"spot"`
	Text        string   `xml:",chardata"`
	App         string   `xml:"app"`
	StationName string   `xml:"StationName"`
	Dxcall      string   `xml:"dxcall"`
	Frequency   string   `xml:"frequency"`
	Spottercall string   `xml:"spottercall"`
	Timestamp   string   `xml:"timestamp"`
	Action      string   `xml:"action"`
	Mode        string   `xml:"mode"`
	Comment     string   `xml:"comment"`
	Status      string   `xml:"status"`
	Statuslist  string   `xml:"statuslist"`
}

type Dynamicresults struct {
	XMLName xml.Name `xml:"dynamicresults"`
	Text    string   `xml:",chardata"`
	Contest string   `xml:"contest"`
	Call    string   `xml:"call"`
	Ops     string   `xml:"ops"`
	Class   struct {
		Text        string `xml:",chardata"`
		Power       string `xml:"power,attr"`
		Assisted    string `xml:"assisted,attr"`
		Transmitter string `xml:"transmitter,attr"`
		Ops         string `xml:"ops,attr"`
		Bands       string `xml:"bands,attr"`
		Mode        string `xml:"mode,attr"`
		Overlay     string `xml:"overlay,attr"`
	} `xml:"class"`
	Club string `xml:"club"`
	Qth  struct {
		Text        string `xml:",chardata"`
		Dxcccountry string `xml:"dxcccountry"`
		Cqzone      string `xml:"cqzone"`
		Iaruzone    string `xml:"iaruzone"`
		Arrlsection string `xml:"arrlsection"`
		Stprvoth    string `xml:"stprvoth"`
		Grid6       string `xml:"grid6"`
	} `xml:"qth"`
	Breakdown struct {
		Text string `xml:",chardata"`
		Qso  []struct {
			Text string `xml:",chardata"`
			Band string `xml:"band,attr"`
			Mode string `xml:"mode,attr"`
		} `xml:"qso"`
		Point []struct {
			Text string `xml:",chardata"`
			Band string `xml:"band,attr"`
			Mode string `xml:"mode,attr"`
		} `xml:"point"`
	} `xml:"breakdown"`
	Score     string `xml:"score"`
	Timestamp string `xml:"timestamp"`
}

type Spectrum struct {
	XMLName            xml.Name `xml:"Spectrum"`
	Text               string   `xml:",chardata"`
	App                string   `xml:"app"`
	Name               string   `xml:"Name"`
	LowScopeFrequency  string   `xml:"LowScopeFrequency"`
	HighScopeFrequency string   `xml:"HighScopeFrequency"`
	ScalingFactor      string   `xml:"ScalingFactor"`
	DataCount          string   `xml:"DataCount"`
	SpectrumData       string   `xml:"SpectrumData"`
}

type RadioSetfrequency struct {
	XMLName     xml.Name `xml:"radio_setfrequency"`
	Text        string   `xml:",chardata"`
	App         string   `xml:"app"`
	Radionr     string   `xml:"radionr"`
	Frequency   string   `xml:"frequency"`
	Mousebutton string   `xml:"mousebutton"`
}

type N1MMRotorTurn struct {
	XMLName       xml.Name `xml:"N1MMRotor"`
	Text          string   `xml:",chardata"`
	Rotor         string   `xml:"rotor"`
	Goazi         string   `xml:"goazi"`
	Offset        string   `xml:"offset"`
	Bidirectional string   `xml:"bidirectional"`
	Freqband      string   `xml:"freqband"`
}

type N1MMRotorStop struct {
	XMLName xml.Name `xml:"N1MMRotor"`
	Text    string   `xml:",chardata"`
	Stop    struct {
		Text     string `xml:",chardata"`
		Rotor    string `xml:"rotor"`
		Freqband string `xml:"freqband"`
	} `xml:"stop"`
}
