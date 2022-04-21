package domains

import "encoding/xml"

type ResponseDial struct {
	XMLName xml.Name `xml:"Response"`
	Dial    Dial     `xml:"Dial,omitempty"`
	Hangup  Hangup   `xml:"Hangup,omitempty"`
}

type ResponseDialNumber struct {
	XMLName    xml.Name   `xml:"Response"`
	DialNumber DialNumber `xml:"Dial,omitempty"`
	Hangup     Hangup     `xml:"Hangup,omitempty"`
}

type ResponseConference struct {
	XMLName        xml.Name       `xml:"Response"`
	DialConference DialConference `xml:"Dial,omitempty"`
	Pause          Pause          `xml:"Pause,omitempty"`
	Hangup         Hangup         `xml:"Hangup,omitempty"`
}

type Dial struct {
	Value       string `xml:",chardata"`
	CallBackURL string `xml:"callbackUrl,attr"`
}

type DialNumber struct {
	XMLName xml.Name `xml:"Dial"`
	Number  Number   `xml:"Number,omitempty"`
}

type DialConference struct {
	XMLName    xml.Name   `xml:"Dial"`
	Conference Conference `xml:"Conference,omitempty"`
}

type Number struct {
	Value      string `xml:",chardata"`
	SendDigits string `xml:"sendDigits,attr"`
}
