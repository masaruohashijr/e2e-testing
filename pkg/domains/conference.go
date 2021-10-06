package domains

type Conference struct {
	Value                  string          `xml:",chardata"`
	Sid                    string          `json:"sid,omitempty"`
	FriendlyName           string          `json:"friendly_name,omitempty"`
	AccountSid             string          `json:"account_sid,omitempty"`
	ActiveParticipants     string          `json:"active_participants_count,omitempty"`
	Status                 string          `json:"status,omitempty"`
	Uri                    string          `json:"uri,omitempty"`
	DateCreated            string          `json:"date_created,omitempty"`
	DateUpdated            string          `json:"date_updated,omitempty"`
	StartConferenceOnEnter bool            `xml:"startConferenceOnEnter,attr"`
	CallbackUrl            string          `xml:"callbackUrl,attr"`
	HangupOnStar           bool            `xml:"hangupOnStar,attr"`
	MaxParticipants        int             `xml:"maxParticipants,attr"`
	SubresourceUris        SubresourceUris `json:"subresource_uris,omitempty"`
}

type SubresourceUris struct {
	Participants []Participant
}

type Participant struct {
	ConferenceSid string `json:"conference_sid"`
	Body          string `json:"body"`
	Sid           string `json:"sid"`
	Muted         string `json:"muted"`
	Deaf          string `json:"deaf"`
	CallerName    string `json:"caller_name"`
	CallerNumber  string `json:"caller_number"`
	Duration      string `json:"duration"`
	DateUpdated   string `json:"date_updated"`
	DateCreated   string `json:"date_created"`
	Price         string `json:"price"`
	Uri           string `json:"uri"`
	ApiVersion    string `json:"api_version"`
	AccountSid    string `json:"account_sid"`
}
