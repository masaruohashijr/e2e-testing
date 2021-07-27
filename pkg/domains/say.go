package domains

import "encoding/xml"

type ResponseSay struct {
	XMLName xml.Name `xml:"Response"`
	Pause   Pause    `xml:"Pause"`
	Say     Say      `xml:"Say"`
}

type Say struct {
	Value    string `xml:",chardata"`
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
}
