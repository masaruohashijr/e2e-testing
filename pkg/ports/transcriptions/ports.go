package transcriptions

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	TranscribeRecording(recordingSid string) (domains.Transcription, error)
	TranscribeAudioUrl(audioUrl string) (domains.Transcription, error)
	ViewTranscription(transcriptionSid string) (domains.Transcription, error)
	ListTranscriptions() ([]domains.Transcription, error)
}

type SecondaryPort interface {
	TranscribeRecording(recordingSid string) (domains.Transcription, error)
	TranscribeAudioUrl(audioUrl string) (domains.Transcription, error)
	ViewTranscription(transcriptionSid string) (domains.Transcription, error)
	ListTranscriptions() ([]domains.Transcription, error)
}
