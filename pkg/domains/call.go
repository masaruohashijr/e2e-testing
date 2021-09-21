package domains

type CallsResponse struct {
	FirstPageUri    string `json:"first_page_uri,omitempty"`
	End             string `json:"end,omitempty"`
	Total           string `json:"total,omitempty"`
	PreviousPageUri string `json:"previous_page_uri,omitempty"`
	NumPages        string `json:"num_pages,omitempty"`
	Calls           []Call `json:"calls,omitempty"`
}

type Call struct {
	DateUpdated     string `json:"date_updated,omitempty"`
	ParentCallSid   string `json:"parent_call_sid,omitempty"`
	Duration        int    `json:"duration,omitempty"`
	From            string `json:"from,omitempty"`
	To              string `json:"to,omitempty"`
	CallerIdBlocked string `json:"caller_id_blocked,omitempty"`
	AnsweredBy      string `json:"answered_by,omitempty"`
	Sid             string `json:"sid,omitempty"`
	RecordingsCount string `json:"recordings_count,omitempty"`
	Price           string `json:"price,omitempty"`
	ApiVersion      string `json:"api_version,omitempty"`
	Status          string `json:"status,omitempty"`
	Direction       string `json:"direction,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
	DateCreated     string `json:"date_created,omitempty"`
	ForwardedFrom   string `json:"forwarded_from,omitempty"`
	AccountSid      string `json:"account_sid,omitempty"`
	DurationBilled  int    `json:"duration_billed,omitempty"`
	PhoneNumberSid  string `json:"phone_number_sid,omitempty"`
}
