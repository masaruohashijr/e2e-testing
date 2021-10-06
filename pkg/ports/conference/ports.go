package conference

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	ViewConference(conferenceSid string) (domains.Conference, error)
	ListConferences(friendlyName string) ([]domains.Conference, error)
	ViewParticipant(conferenceSid, participantSid string) (domains.Participant, error)
	ListParticipants(conferenceSid string) ([]domains.Participant, error)
	MuteDeafParticipant(conferenceSid, participantSid string) (domains.Participant, error)
	PlayAudioToParticipant(conferenceSid, participantSid string) (domains.Participant, error)
	HangupParticipant(conferenceSid, participantSid string) (domains.Participant, error)
}

type SecondaryPort interface {
	ViewConference(conferenceSid string) (domains.Conference, error)
	ListConferences(friendlyName string) ([]domains.Conference, error)
	ViewParticipant(conferenceSid, participantSid string) (domains.Participant, error)
	ListParticipants(conferenceSid string) ([]domains.Participant, error)
	MuteDeafParticipant(conferenceSid, participantSid string) (domains.Participant, error)
	PlayAudioToParticipant(conferenceSid, participantSid string) (domains.Participant, error)
	HangupParticipant(conferenceSid, participantSid string) (domains.Participant, error)
}
