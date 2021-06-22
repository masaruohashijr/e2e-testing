Feature: record    
  In order to record a call
  As an end user
  I want to try to make a call and record this call until key "#" is pressed.

  Scenario: Pause a sequence of sentences

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberA" configured to say "You would never break the chain." #-> configuredToSay(numberB string, text string)
      And "NumberB" configured to record calls #-> configuredToRecordCalls(numberB string)
      When I make a call from "NumberA" to "NumbeB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberA" should get transcription "You would never break the chain." #-> shouldGetTranscription(numberA string, transcription string)