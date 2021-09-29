Feature: Transcribe_Audio_Url
  In order to transcribe an audio
  As an end user
  I want to provide an audio and view its transcription.

  Scenario: Transcribe Audio Url

    Given my test setup runs
      And I provide an audio url "http://mohashi.ngrok.io/Audio"
      When I transcribe audio url "http://mohashi.ngrok.io/Audio"
      Then I should get transcription text as "what we do in life echoes in eternity"
      