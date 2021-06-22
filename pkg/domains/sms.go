package domains

import "encoding/xml"

type ResponseSMS struct {
	XMLName xml.Name `xml:"Response"`
	Sms     Sms      `xml:"Sms,omitempty"`
}

type Sms struct {
	From           string `xml:"from,attr"`
	To             string `xml:"to,attr"`
	Value          string `xml:",chardata"`
	StatusCallback string `xml:"statusCallback,attr"`
}
