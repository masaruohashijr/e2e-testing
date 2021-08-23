package domains

import "encoding/xml"

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

type ResponseIncomingPhoneNumber struct {
	XMLName             xml.Name `xml:"Response"`
	IncomingPhoneNumber IncomingPhoneNumber
}

type IncomingPhoneNumber struct {
	DateUpdated          string       `json:"date_updated" xml:"DateUpdated"`
	VoiceURL             string       `json:"voice_url" xml:"VoiceUrl"`
	VoiceFallbackMethod  string       `json:"voice_fallback_method" xml:"VoiceFallbackMethod"`
	Capabilities         Capabilities `json:"capabilities" xml:"Capabilities"`
	Sid                  string       `json:"sid" xml:"Sid"`
	HeartbeatMethod      string       `json:"heartbeat_method" xml:"HeartbeatMethod"`
	Type                 string       `json:"type" xml:"Type"`
	StatusCallbackMethod string       `json:"status_callback_method" xml:"StatusCallbackMethod"`
	VoiceFallbackURL     string       `json:"voice_fallback_url" xml:"VoiceFallbackURL"`
	PhoneNumber          string       `json:"phone_number" xml:"PhoneNumber"`
	HangupCallback       string       `json:"hangup_callback" xml:"HangupCallback"`
	HangupCallbackMethod string       `json:"hangup_callback_method" xml:"HangupCallbackMethod"`
	HeartbeatURL         string       `json:"heartbeat_url" xml:"HeartbeatURL"`
	MmsURL               string       `json:"sms_url" xml:"MmsURL"`
	VoiceMethod          string       `json:"voice_method" xml:"VoiceMethod"`
	VoiceCallerIDLookup  bool         `json:"voice_caller_id_lookup" xml:"VoiceCallerIDLookup"`
	FriendlyName         string       `json:"friendly_name" xml:"FriendlyName"`
	URI                  string       `json:"uri" xml:"URI"`
	SmsFallbackURL       string       `json:"sms_fallback_url" xml:"SmsFallbackURL"`
	AccountSid           string       `json:"account_sid" xml:"AccountSid"`
	SmsMethod            string       `json:"sms_method" xml:"SmsMethod"`
	NextRenewalDate      string       `json:"next_renewal_date" xml:"NextRenewalDate"`
	DateCreated          string       `json:"date_created" xml:"DateCreated"`
	StatusCallback       string       `json:"status_callback" xml:"StatusCallback"`
}

type Capabilities struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"sms"`
	MMS   bool `json:"mms"`
}
