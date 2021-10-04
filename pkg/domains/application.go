package domains

type ResponseApplication struct {
	Page            string        `json:"page,omitempty"`
	NumPages        string        `json:"num_pages,omitempty"`
	PageSize        string        `json:"page_size,omitempty"`
	Total           string        `json:"total,omitempty"`
	Start           string        `json:"start,omitempty"`
	End             string        `json:"end,omitempty"`
	Uri             string        `json:"uri,omitempty"`
	FirstPageUri    string        `json:"first_page_uri,omitempty"`
	PreviousPageUri string        `json:"previous_page_uri,omitempty"`
	NextPageUri     string        `json:"next_page_uri,omitempty"`
	LastPageUri     string        `json:"last_page_uri,omitempty"`
	Applications    []Application `json:"applications,omitempty"`
}

type Application struct {
	Sid                  string `json:"sid,omitempty"`
	DateCreated          string `json:"date_created,omitempty"`
	DateUpdated          string `json:"date_updated,omitempty"`
	AccountSid           string `json:"account_sid,omitempty"`
	ClientCount          string `json:"client_count,omitempty"`
	FriendlyName         string `json:"friendly_name,omitempty"`
	ApiVersion           string `json:"api_version,omitempty"`
	VoiceUrl             string `json:"voice_url,omitempty"`
	Status               string `json:"status,omitempty"`
	Type                 string `json:"type,omitempty"`
	AudioUrl             string `json:"audio_url,omitempty"`
	Duration             string `json:"duration,omitempty"`
	TranscriptionText    string `json:"transcription_text,omitempty"`
	VoiceMethod          string `json:"voice_method,omitempty"`
	VoiceFallbackUrl     string `json:"voice_fallback_url,omitempty"`
	VoiceFallbackMethod  string `json:"voice_fallback_method,omitempty"`
	HeartbeatUrl         string `json:"heartbeat_url,omitempty"`
	Heartbeat_method     string `json:"heartbeat_method,omitempty"`
	StatusCallback       string `json:"status_callback,omitempty"`
	StatusCallbackMethod string `json:"status_callback_method,omitempty"`
	HangupCallback       string `json:"hangup_callback,omitempty"`
	HangupCallbackMethod string `json:"hangup_callback_method,omitempty"`
	VoiceCallerIdLookup  string `json:"voice_caller_id_lookup,omitempty"`
	SmsUrl               string `json:"sms_url,omitempty"`
	SmsMethod            string `json:"sms_method,omitempty"`
	SmsFallbackUrl       string `json:"sms_fallback_url,omitempty"`
	SmsFallbackMethod    string `json:"sms_fallback_method,omitempty"`
	SmsStatusCallback    string `json:"sms_status_callback,omitempty"`
	SmsStatusCallbackUrl string `json:"sms_status_callback_url,omitempty"`
	Uri                  string `json:"uri,omitempty"`
}
