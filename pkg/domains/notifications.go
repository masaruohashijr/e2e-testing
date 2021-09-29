package domains

import "encoding/xml"

type ResponseNotifications struct {
	XMLName xml.Name `xml:"Response"`
	Mms     Mms      `xml:"Mms,omitempty"`
}

type Notification struct {
	From           string `xml:"from,attr"`
	To             string `xml:"to,attr"`
	Value          string `xml:",chardata"`
	MediaUrl       string `xml:"mediaUrl,attr"`
	StatusCallback string `xml:"statusCallback,attr"`
}
