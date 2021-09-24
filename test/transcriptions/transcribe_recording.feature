
Feature: Transcribe Recording
  In order to transcribe a recording
  As an end user
  I want to record a call and view the recording transcription.

  Scenario: Transcribe Recording

    Given my test setup runs
      And "NumberA" configured to pause 3 seconds
      And Append To "NumberA" config say "what we do in life echoes in eternity"
      When I make a call from "NumberA" to "NumberB"
      And I record current call from "NumberA" to "NumberB" for 5 seconds
      And I transcribe last recoding from "NumberA" to "NumberB"
      Then I should get last transcription text as "what we do in life echoes in eternity"