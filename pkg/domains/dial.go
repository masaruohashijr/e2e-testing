package domains

import "encoding/xml"

type ResponseDial struct {
	XMLName xml.Name `xml:"Response"`
	Dial    Dial     `xml:"Dial,omitempty"`
	Hangup  Hangup   `xml:"Hangup,omitempty"`
}

type Dial struct {
	Value       string `xml:",chardata"`
	CallBackURL string `xml:"callbackUrl,attr"`
}
