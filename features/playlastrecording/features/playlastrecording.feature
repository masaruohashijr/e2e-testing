Feature: play last recording    
  In order to listen the last recording of my calls
  As an end user
  I want to call a number and set to play the last recording.

  Scenario: Play Last Recording

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberA" configured to say "This is the last recording" #-> configuredToSay(numberB string, text string)
      And "NumberB" configured to record calls #-> configuredToRecordCalls(numberB string)
      When I make a call from "NumberA" to "NumbeB" #-> iMakeACallFromTo(numberA string, numberB string)
      And "NumberB" configured to play last recording #-> configuredToPlayLastRecording(numberB string)
      Then "NumberA" should get transcription "This is the last recording" #-> shouldGetTranscription(numberA string, transcription string)   