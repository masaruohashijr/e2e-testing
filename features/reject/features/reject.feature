Feature: reject    
  In order to reject a call
  As an end user
  I want to try to make a call to a Number and listen to a not-in-service message.

  Scenario: Pause a sequence of sentences

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberB" configured to dial to "NumberC" #-> configuredToDialTo(numberB string, numberA string)
      And "NumberB" configured to gather speech #-> configuredToGatherSpeech(numberB string, text string)
      And "NumberC" configured to reject with "not-in-service" reason #-> configuredToRejectWithReason(numberC string, reason string)
      When I make a call from "NumberA" to "NumbeB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberB" should get speech "not-in-service" #-> shouldGetSpeech(numberA string, speech string)