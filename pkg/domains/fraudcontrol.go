package domains

type ResponseFraudControl struct {
	Frauds          []Fraud `json:"frauds,omitempty"`
	Page            string  `json:"page,omitempty"`
	PageSize        string  `json:"page_size,omitempty"`
	NumPages        string  `json:"num_pages,omitempty"`
	Start           string  `json:"start,omitempty"`
	Total           string  `json:"total,omitempty"`
	End             string  `json:"end,omitempty"`
	Uri             string  `json:"uri,omitempty"`
	FirstPageUri    string  `json:"first_page_uri,omitempty"`
	LastPageUri     string  `json:"last_page_uri,omitempty"`
	NextPageUri     string  `json:"next_page_uri,omitempty"`
	PreviousPageUri string  `json:"previous_page_uri,omitempty"`
}

type Fraud struct {
	WhiteListed WhiteListed `json:"whitelisted,omitempty"`
	Blocked     Blocked     `json:"blocked,omitempty"`
}

type Blocked struct {
	DateUpdated   string `json:"date_updated,omitempty"`
	DateCreated   string `json:"date_created,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	Sid           string `json:"sid,omitempty"`
	IsLock        string `json:"is_lock,omitempty"`
	CountryName   string `json:"country_name,omitempty"`
	CountryPrefix string `json:"country_prefix,omitempty"`
	SmsEnabled    string `json:"sms_enabled,omitempty"`
}

type Authorized struct {
	DateUpdated     string `json:"date_updated"`
	SmsEnabled      string `json:"sms_enabled,omitempty"`
	CountryCode     string `json:"country_code,omitempty"`
	IsLock          string `json:"is_lock,omitempty"`
	ExpirationDate  string `json:"expiration_date,omitempty"`
	MobileEnabled   string `json:"mobile_enabled,omitempty"`
	LandlineEnabled string `json:"landline_enabled,omitempty"`
	Sid             string `json:"sid,omitempty"`
	CountryName     string `json:"country_name,omitempty"`
	DateCreated     string `json:"date_created,omitempty"`
	CountryPrefix   string `json:"country_prefix,omitempty"`
}

type WhiteListed struct {
	IsLock          string `json:"is_lock,omitempty"`
	MobileEnabled   string `json:"mobile_enabled,omitempty"`
	LandlineEnabled string `json:"landline_enabled,omitempty"`
	DateUpdated     string `json:"date_updated"`
	CountryCode     string `json:"country_code,omitempty"`
	Sid             string `json:"sid,omitempty"`
	CountryName     string `json:"country_name,omitempty"`
	DateCreated     string `json:"date_created,omitempty"`
	SmsEnabled      string `json:"sms_enabled,omitempty"`
	CountryPrefix   string `json:"country_prefix,omitempty"`
}
