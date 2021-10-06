package domains

import "encoding/xml"

type ResponseCarrierLookup struct {
	XMLName         xml.Name        `xml:"Response"`
	CarrierLookups  []CarrierLookup `xml:"CarrierLookups" json:"carrier_lookups"`
	Page            string          `json:"page,omitempty"`
	PageSize        string          `json:"page_size,omitempty"`
	NumPages        string          `json:"num_pages,omitempty"`
	Start           string          `json:"start,omitempty"`
	Total           string          `json:"total,omitempty"`
	End             string          `json:"end,omitempty"`
	Uri             string          `json:"uri,omitempty"`
	FirstPageUri    string          `json:"first_page_uri,omitempty"`
	LastPageUri     string          `json:"last_page_uri,omitempty"`
	NextPageUri     string          `json:"next_page_uri,omitempty"`
	PreviousPageUri string          `json:"previous_page_uri,omitempty"`
}

type ResponseCNAM struct {
	CnamLookups     []CNAM `json:"cnam_lookups"`
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
}

type ResponseBNA struct {
	BNALookups      []BNA  `json:"bna_lookups"`
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
}

type CarrierLookup struct {
	PhoneNumber string `xml:"PhoneNumber" json:"phone_number,omitempty"`
	DateUpdated string `xml:"DateUpdated" json:"date_updated,omitempty"`
	Price       string `xml:"Price" json:"price,omitempty"`
	CarrierId   string `xml:"CarrierId" json:"carrier_id,omitempty"`
	CountryCode string `xml:"CountryCode" json:"country_code,omitempty"`
	Network     string `xml:"Network" json:"network,omitempty"`
	Mobile      string `xml:"Mobile" json:"mobile,omitempty"`
	Uri         string `xml:"Uri" json:"uri,omitempty"`
	AccountSid  string `xml:"AccountSid" json:"account_sid,omitempty"`
	Sid         string `xml:"Sid" json:"sid,omitempty"`
	DateCreated string `xml:"DateCreated" json:"date_created,omitempty"`
	ApiVersion  string `xml:"ApiVersion" json:"api_version,omitempty"`
	Mcc         string `xml:"Mcc" json:"mcc,omitempty"`
	Mnc         string `xml:"Mnc" json:"mnc,omitempty"`
}

type CNAM struct {
	PhoneNumber string `json:"phone_number"`
	Body        string `json:"body"`
	Sid         string `json:"sid"`
	DateUpdated string `json:"date_updated"`
	DateCreated string `json:"date_created"`
	Price       string `json:"price"`
	Uri         string `json:"uri"`
	ApiVersion  string `json:"api_version"`
	AccountSid  string `json:"account_sid"`
}

type BNA struct {
	PhoneNumber string `json:"phone_number"`
	Sid         string `json:"sid"`
	DateUpdated string `json:"date_updated"`
	DateCreated string `json:"date_created"`
	Price       string `json:"price"`
	Uri         string `json:"uri"`
	ApiVersion  string `json:"api_version"`
	AccountSid  string `json:"account_sid"`
	City        string `json:"city"`
	State       string `json:"state"`
	CountryCode string `json:"country_code"`
}
