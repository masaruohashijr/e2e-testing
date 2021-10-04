package domains

type ResponseNotification struct {
	Page            string         `json:"page,omitempty"`
	NumPages        string         `json:"num_pages,omitempty"`
	PageSize        string         `json:"page_size,omitempty"`
	Total           string         `json:"total,omitempty"`
	Start           string         `json:"start,omitempty"`
	End             string         `json:"end,omitempty"`
	Uri             string         `json:"uri,omitempty"`
	FirstPageUri    string         `json:"first_page_uri,omitempty"`
	PreviousPageUri string         `json:"previous_page_uri,omitempty"`
	NextPageUri     string         `json:"next_page_uri,omitempty"`
	LastPageUri     string         `json:"last_page_uri,omitempty"`
	Notifications   []Notification `json:"notifications,omitempty"`
}

type Notification struct {
	Sid              string `json:"sid,omitempty"`
	AccountSid       string `json:"account_sid,omitempty"`
	CallSid          string `json:"call_sid,omitempty"`
	ApiVersion       string `json:"api_version,omitempty"`
	Duration         string `json:"duration,omitempty"`
	DateCreated      string `json:"date_created,omitempty"`
	DateUpdated      string `json:"date_updated,omitempty"`
	Log              string `json:"log,omitempty"`
	ErrorCode        string `json:"error_code,omitempty"`
	MoreInfo         string `json:"more_info,omitempty"`
	MessageText      string `json:"message_text,omitempty"`
	MessageDate      string `json:"message_date,omitempty"`
	ResponseBody     string `json:"response_body,omitempty"`
	RequestMethod    string `json:"request_method,omitempty"`
	RequestUrl       string `json:"request_url,omitempty"`
	RequestVariables string `json:"request_variables,omitempty"`
	RequestHeaders   string `json:"response_headers,omitempty"`
	Uri              string `json:"uri,omitempty"`
}
