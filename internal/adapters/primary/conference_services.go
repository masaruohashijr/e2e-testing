package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/conference"
)

type port_conference struct {
	driven conference.SecondaryPort
}

func NewConferenceService(driven conference.SecondaryPort) conference.PrimaryPort {
	return &port_conference{
		driven,
	}
}

func (p *port_conference) ViewConference(conferenceSid string) (domains.Conference, error) {
	conference, err := p.driven.ViewConference(conferenceSid)
	return conference, err
}

func (p *port_conference) ListConferences(friendlyName string) ([]domains.Conference, error) {
	conferences, err := p.driven.ListConferences(friendlyName)
	return conferences, err
}

func (p *port_conference) ViewParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	participant, err := p.driven.ViewParticipant(conferenceSid, participantSid)
	return participant, err
}

func (p *port_conference) ListParticipants(conferenceSid string) ([]domains.Participant, error) {
	participants, err := p.driven.ListParticipants(conferenceSid)
	return participants, err
}
func (p *port_conference) MuteDeafParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	participant, err := p.driven.MuteDeafParticipant(conferenceSid, participantSid)
	return participant, err
}
func (p *port_conference) PlayAudioToParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	participant, err := p.driven.PlayAudioToParticipant(conferenceSid, participantSid)
	return participant, err
}
func (p *port_conference) HangupParticipant(conferenceSid, participantSid string) (domains.Participant, error) {
	participant, err := p.driven.HangupParticipant(conferenceSid, participantSid)
	return participant, err
}
