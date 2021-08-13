package domains

type AvailablePhoneNumbersPageResponse struct {
	AvailablePhoneNumbers []PageAvailablePhoneNumberResponse `json:"available_phone_numbers"`
	URI                   string                             `json:"uri"`
}

type IncomingPhoneNumbersPageResponse struct {
	IncomingPhoneNumbers []PageIncomingPhoneNumberResponse `json:"incoming_phone_numbers"`
	URI                  string                            `json:"uri"`
}

type PageAvailablePhoneNumberResponse struct {
	AddressRequirements string                                       `json:"address_requirements"`
	Beta                bool                                         `json:"beta"`
	Capabilities        PageAvailablePhoneNumberCapabilitiesResponse `json:"capabilities"`
	FriendlyName        string                                       `json:"friendly_name"`
	IsoCountry          string                                       `json:"iso_country"`
	Lata                *string                                      `json:"lata,omitempty"`
	Latitude            string                                       `json:"latitude"`
	Locality            *string                                      `json:"locality,omitempty"`
	Longitude           string                                       `json:"longitude"`
	PhoneNumber         string                                       `json:"phone_number"`
	PostalCode          *string                                      `json:"postal_code,omitempty"`
	RateCenter          *string                                      `json:"rate_center,omitempty"`
	Region              *string                                      `json:"region,omitempty"`
}

type PageAvailablePhoneNumberCapabilitiesResponse struct {
	Fax   *bool `json:"fax,omitempty"`
	Mms   bool  `json:"MMS"`
	Sms   bool  `json:"SMS"`
	Voice bool  `json:"voice"`
}

type PageIncomingPhoneNumberResponse struct {
	AddressRequirements string  `json:"address_requirements"`
	Beta                bool    `json:"beta"`
	FriendlyName        string  `json:"friendly_name"`
	IsoCountry          string  `json:"iso_country"`
	Lata                *string `json:"lata,omitempty"`
	Latitude            string  `json:"latitude"`
	Locality            *string `json:"locality,omitempty"`
	Longitude           string  `json:"longitude"`
	PhoneNumber         string  `json:"phone_number"`
	PostalCode          *string `json:"postal_code,omitempty"`
	RateCenter          *string `json:"rate_center,omitempty"`
	Region              *string `json:"region,omitempty"`
}

type IncomingPhoneNumber struct {
	DateUpdated          string       `json:"date_updated"`
	VoiceURL             string       `json:"voice_url"`
	VoiceFallbackMethod  string       `json:"voice_fallback_method"`
	Capabilities         Capabilities `json:"capabilities"`
	Sid                  string       `json:"sid"`
	HeartbeatMethod      string       `json:"heartbeat_method"`
	Type                 string       `json:"type"`
	StatusCallbackMethod string       `json:"status_callback_method"`
	VoiceFallbackURL     string       `json:"voice_fallback_url"`
	PhoneNumber          string       `json:"phone_number"`
	HangupCallback       string       `json:"hangup_callback"`
	HangupCallbackMethod string       `json:"hangup_callback_method"`
	HeartbeatURL         string       `json:"heartbeat_url"`
	SmsURL               string       `json:"sms_url"`
	VoiceMethod          string       `json:"voice_method"`
	VoiceCallerIDLookup  bool         `json:"voice_caller_id_lookup"`
	FriendlyName         string       `json:"friendly_name"`
	URI                  string       `json:"uri"`
	SmsFallbackURL       string       `json:"sms_fallback_url"`
	AccountSid           string       `json:"account_sid"`
	SmsMethod            string       `json:"sms_method"`
	NextRenewalDate      string       `json:"next_renewal_date"`
	DateCreated          string       `json:"date_created"`
	StatusCallback       string       `json:"status_callback"`
}

type Capabilities struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"sms"`
	MMS   bool `json:"mms"`
}
