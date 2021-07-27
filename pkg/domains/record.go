package domains

import "encoding/xml"

type ResponseRecord struct {
	XMLName xml.Name `xml:"Response"`
	Pause   Pause    `xml:"Pause,omitempty"`
	Record  Record   `xml:"Record,omitempty"`
	Hangup  Hangup   `xml:"Hangup,omitempty"`
}

type Record struct {
	Background         bool   `xml:"background,attr"`
	Action             string `xml:"action,attr"`
	MaxLength          int    `xml:"maxLength,attr"`
	FileFormat         string `xml:"fileFormat,attr"`
	Transcribe         bool   `xml:"transcribe,attr"`
	TranscribeCallback string `xml:"transcribeCallback,attr"`
}
