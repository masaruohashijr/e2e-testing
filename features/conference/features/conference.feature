Feature: conference    
  In order to make a conference call
  As an end user
  I want to dial to a number, create and set a conference room

  Scenario: Make a Conference

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberB" configured as conference with size 3 #-> configuredAsConferenceWithSize(numberB string, size int)
      And "NumberC" configured to say "What a wonderful world" #-> configuredToSay(numberC string, message string)
      And "NumberA" configured to gather speech #-> configuredToGatherSpeech(numberB string, size int)
      When I make a call from "NumberC" to "NumberB" #-> iMakeACallFromTo(numberC string, numberB string) 
      And  I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberA string, numberB string)  
      Then "NumberA" should get speech "What a wonderful world" #-> shouldGetSpeech(numberA string, text string)   

