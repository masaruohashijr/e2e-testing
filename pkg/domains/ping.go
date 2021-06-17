package domains

import "encoding/xml"

type ResponsePing struct {
	XMLName xml.Name `xml:"Response"`
	Ping    Ping     `xml:"Ping,omitempty"`
}

type Ping struct {
	Value string `xml:",chardata"`
}
