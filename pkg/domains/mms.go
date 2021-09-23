package domains

import "encoding/xml"

type ResponseMMS struct {
	XMLName xml.Name `xml:"Response"`
	Mms     Mms      `xml:"Mms,omitempty"`
}

type MMSResponse struct {
	Page            string `json:"page,omitempty"`
	PageSize        string `json:"page_size,omitempty"`
	NumPages        string `json:"num_pages,omitempty"`
	Start           string `json:"start,omitempty"`
	Total           string `json:"total,omitempty"`
	End             string `json:"end,omitempty"`
	Uri             string `json:"uri,omitempty"`
	FirstPageUri    string `json:"first_page_uri,omitempty"`
	LastPageUri     string `json:"last_page_uri,omitempty"`
	NextPageUri     string `json:"next_page_uri,omitempty"`
	PreviousPageUri string `json:"previous_page_uri,omitempty"`
	MmsMessages     []Mms  `json:"mms_messages,omitempty"`
}

type Mms struct {
	AccountSid     string `json:"account_sid,omitempty"`
	ApiVersion     string `json:"api_version,omitempty"`
	Value          string `xml:",chardata"`
	Body           string `json:"body,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
	DateSent       string `json:"date_sent,omitempty"`
	DateUpdated    string `json:"date_updated,omitempty"`
	Direction      string `json:"direction,omitempty"`
	From           string `xml:"from,attr" json:"from,omitempty"`
	MediaUrl       string `xml:"mediaUrl,attr" json:"media_url,omitempty"`
	Sid            string `json:"mms_sid,omitempty"`
	Price          string `json:"price,omitempty"`
	Status         string `json:"status,omitempty"`
	StatusCallback string `xml:"statusCallback,attr"`
	To             string `xml:"to,attr" json:"to,omitempty"`
	Uri            string `json:"uri,omitempty"`
}
