package domains

import "encoding/xml"

type ResponseSMS struct {
	XMLName xml.Name `xml:"Response"`
	Sms     Sms      `xml:"Sms,omitempty"`
}

type SMSResponse struct {
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
	SmsMessages     []Sms  `json:"sms_messages,omitempty"`
}

type Sms struct {
	StatusCallback string `xml:"statusCallback,attr"`
	Value          string `xml:",chardata"`
	From           string `xml:"from,attr" json:"from,omitempty"`
	To             string `xml:"to,attr" json:"tp,omitempty"`
	Body           string `json:"body,omitempty"`
	Status         string `json:"status,omitempty"`
	Direction      string `json:"direction,omitempty"`
	DateUpdated    string `json:"date_updated,omitempty"`
	Price          string `json:"price,omitempty"`
	Uri            string `json:"uri,omitempty"`
	AccountSid     string `json:"account_sid,omitempty"`
	Sid            string `json:"sid,omitempty"`
	DateSent       string `json:"date_sent,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
	ApiVersion     string `json:"api_version,omitempty"`
}
