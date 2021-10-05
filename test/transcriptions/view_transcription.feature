Feature: View_Transcription
  In order to record a call
  As an end user
  I want to try to make a call and record this call and get transcription of it.

  Scenario: View Transcrition

    Given my test setup runs
      And "NumberB" configured to pause 3 seconds
      And Append To "NumberB" config say "what we do in life echoes in eternity"
      When I record a call from "NumberA" to "NumberB" for 8 seconds
      And I transcribe last recording
      Then I should get last transcription text as "what we do in life echoes in eternity"
