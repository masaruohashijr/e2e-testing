Feature: sms    
  In order to my sent SMS messages can be seem by my destination numbers
  And getting a receiving confirmation from them
  As an end user
  I want to send an SMS message

  Scenario: Send an SMS

    Given my test setup runs #-> myTestSetupRuns()
    And created SMS with text "I think to myself"  #-> createdSmsWithText(text string)
    When I send SMS from "NumberA" to "NumberB" #-> iSendSmsFromTo(numberA string, numberB string)
    Then "NumberB" should get SMS with text "I think to myself" #-> shouldGetSmsWithText(text string)