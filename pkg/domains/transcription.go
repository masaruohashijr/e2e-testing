package domains

type ResponseTranscription struct {
	Page            string          `json:"page,omitempty"`
	NumPages        string          `json:"num_pages,omitempty"`
	PageSize        string          `json:"page_size,omitempty"`
	Total           string          `json:"total,omitempty"`
	Start           string          `json:"start,omitempty"`
	End             string          `json:"end,omitempty"`
	Uri             string          `json:"uri,omitempty"`
	FirstPageUri    string          `json:"first_page_uri,omitempty"`
	PreviousPageUri string          `json:"previous_page_uri,omitempty"`
	NextPageUri     string          `json:"next_page_uri,omitempty"`
	LastPageUri     string          `json:"last_page_uri,omitempty"`
	Transcriptions  []Transcription `json:"transcriptions,omitempty"`
}

type Transcription struct {
	Sid                string `json:"sid,omitempty"`
	DateCreated        string `json:"date_created,omitempty"`
	DateUpdated        string `json:"date_updated,omitempty"`
	AccountSid         string `json:"account_sid,omitempty"`
	Status             string `json:"status,omitempty"`
	Type               string `json:"type,omitempty"`
	AudioUrl           string `json:"audio_url,omitempty"`
	Duration           string `json:"duration,omitempty"`
	TranscriptionText  string `json:"transcription_text,omitempty"`
	ApiVersion         string `json:"api_version,omitempty"`
	Price              string `json:"price,omitempty"`
	TranscribeCallback string `json:"transcribe_callback,omitempty"`
	CallbackMethod     string `json:"callback_method,omitempty"`
	Uri                string `json:"uri,omitempty"`
}
