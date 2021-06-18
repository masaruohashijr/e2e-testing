package domains

import "encoding/xml"

type ResponseDial struct {
	XMLName xml.Name `xml:"Response"`
	Dial    Dial     `xml:"Dial,omitempty"`
	Hangup  Hangup   `xml:"Pause,omitempty"`
}

type Dial struct {
	CallBackURL string `xml:"callbackUrl,attr"`
	Value       string `xml:",chardata"`
}
