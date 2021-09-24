
Feature: List_Transcriptions
  In order to list transcriptions
  As an end user
  I want to transcribe a recording and list at least 1 transcription.

  Scenario: List Transcriptions

    Given my test setup runs
      And "NumberA" configured to pause 3 seconds
      And Append To "NumberA" config say "what we do in life echoes in eternity"
      When I make a call from "NumberA" to "NumberB"
      And I record current call from "NumberA" to "NumberB" for 5 seconds
      Then I should list at least 1 transcription