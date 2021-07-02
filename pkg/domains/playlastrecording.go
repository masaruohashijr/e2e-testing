package domains

import "encoding/xml"

type ResponsePlayLastRecording struct {
	XMLName           xml.Name          `xml:"Response"`
	PlayLastRecording PlayLastRecording `xml:"PlayLastRecording,omitempty"`
}

type PlayLastRecording struct {
}
