package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/transcriptions"
)

type port_transcription struct {
	driven transcriptions.SecondaryPort
}

func NewTranscriptionService(driven transcriptions.SecondaryPort) transcriptions.PrimaryPort {
	return &port_transcription{
		driven,
	}
}

func (p *port_transcription) ListTranscriptions() ([]domains.Transcription, error) {
	at, err := p.driven.ListTranscriptions()
	return at, err
}

func (p *port_transcription) ViewTranscription(transcriptionSid string) (domains.Transcription, error) {
	vt, err := p.driven.ViewTranscription(transcriptionSid)
	return vt, err
}

func (p *port_transcription) TranscribeRecording(recordingSid string) (domains.Transcription, error) {
	vt, err := p.driven.TranscribeRecording(recordingSid)
	return vt, err
}

func (p *port_transcription) TranscribeAudioUrl(audioUrl string) (domains.Transcription, error) {
	vt, err := p.driven.TranscribeAudioUrl(audioUrl)
	return vt, err
}
