package domains

import "encoding/xml"

type ResponseRedirect struct {
	XMLName  xml.Name `xml:"Response"`
	Redirect Redirect `xml:"Redirect,omitempty"`
}

type Redirect struct {
	Value string `xml:",chardata"`
}
