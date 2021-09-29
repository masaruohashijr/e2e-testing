package domains

type ResponseRecording struct {
	Page            string      `json:"page,omitempty"`
	NumPages        string      `json:"num_pages,omitempty"`
	PageSize        string      `json:"page_size,omitempty"`
	Total           string      `json:"total,omitempty"`
	Start           string      `json:"start,omitempty"`
	End             string      `json:"end,omitempty"`
	Uri             string      `json:"uri,omitempty"`
	FirstPageUri    string      `json:"first_page_uri,omitempty"`
	PreviousPageUri string      `json:"previous_page_uri,omitempty"`
	NextPageUri     string      `json:"next_page_uri,omitempty"`
	LastPageUri     string      `json:"last_page_uri,omitempty"`
	Recordings      []Recording `json:"recordings,omitempty"`
}

type Recording struct {
	Sid          string `json:"sid,omitempty"`
	AccountSid   string `json:"account_sid,omitempty"`
	CallSid      string `json:"call_sid,omitempty"`
	Duration     string `json:"duration,omitempty"`
	DateCreated  string `json:"date_created,omitempty"`
	ApiVersion   string `json:"api_version,omitempty"`
	DateUpdated  string `json:"date_updated,omitempty"`
	RecordingUrl string `json:"recording_url,omitempty"`
	Uri          string `json:"uri,omitempty"`
}
