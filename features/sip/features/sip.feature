Feature: sip    
  In order to make a seamless call transfer to another number using sip
  As an end user
  I want to make an outgoing dial from an already made and current call using sip

  Scenario: Make a Transfer using SIP

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberB" configured to dial "NumberC" #-> configuredToSay(numberA string, message string)
      And my SIP credentials are "SIP address", "username","password" #-> mySipCredentialsAre(sipAddress string, username string, password string)
      And "NumberC" configured to say "I see skies of blue" #-> configuredToSay(numberA string, message string)
      When I make a call from "NumberA" to "NumbeB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberB" should get speech "I see skies of blue" #-> shouldGetSpeech(numberA string, speech string)