package domains

import "encoding/xml"

type ResponseGather struct {
	XMLName xml.Name `xml:"Response"`
	Pause   Pause    `xml:"Pause"`
	Gather  Gather   `xml:"Gather"`
}

type Gather struct {
	Play        string `xml:"Play,omitempty"`
	Input       string `xml:"input,attr,omitempty"`
	Hints       string `xml:"hints,attr,omitempty"`
	Language    string `xml:"language,attr,omitempty"`
	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	Timeout     int    `xml:"timeout,attr,omitempty"`
	FinishOnKey string `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int    `xml:"numDigits,attr,omitempty"`
}
