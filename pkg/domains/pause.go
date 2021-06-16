package domains

import "encoding/xml"

type ResponsePause struct {
	XMLName xml.Name `xml:"Response"`
	Pause   Pause    `xml:"Pause,omitempty"`
}

type Pause struct {
	Length int `xml:"length,attr"`
}

type Hangup struct {
}
