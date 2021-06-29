package domains

import "encoding/xml"

type ResponseReject struct {
	XMLName xml.Name `xml:"Response"`
	Reject  Reject   `xml:"Reject,omitempty"`
}

type Reject struct {
}
