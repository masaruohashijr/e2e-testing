package domains

import "encoding/xml"

type ResponseRecord struct {
	XMLName xml.Name `xml:"Response"`
	Record  Record   `xml:"Record,omitempty"`
}

type Record struct {
	Background bool   `xml:"background,attr"`
	Action     string `xml:"action,attr"`
	MaxLength  int    `xml:"maxLength,attr"`
	FileFormat string `xml:"fileFormat,attr"`
}
