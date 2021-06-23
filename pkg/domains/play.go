package domains

import "encoding/xml"

type ResponsePlay struct {
	XMLName xml.Name `xml:"Response"`
	Play    Play     `xml:"Play,omitempty"`
	Pause   Pause    `xml:"Pause,omitempty"`
}

type Play struct {
	Value string `xml:",chardata"`
	Loop  int    `xml:"loop,attr"`
}
