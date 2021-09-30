Feature: ListCalls  
  As an end user
  I want to make a call and then list my calls
  So that I can list at least 1 call 

  Scenario: List My Calls

    Given my test setup runs    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 3 seconds
    Then I should list at least 1 call

Feature: MakeCall   
  In order to make a call
  As an end user
  I want to make a call from a number to another number

  Scenario: Make a Call

    Given my test setup runs    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 3 seconds
    Then I should get to view a call from "NumberA" to "NumberB" with status "completed"

Feature: ViewLastCall

  Scenario: View Last Call

    Given my test setup runs    
    And "NumberC" configured to pause 2 seconds
    And Append To "NumberC" config hang up
    When I make a call from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should get last call duration greater than or equal to 2 seconds

Feature: ListSMS    
  In order to send an SMS during a call
  As an end user
  I want to make a call and send a text as SMS.

Scenario: List Sms
    Given my test setup runs
    When I send SMS "Actions speak louder than words" from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should list at least 1 SMS from "NumberB" to "NumberC"

Feature: SendSMS    
  In order to send an SMS during a call
  As an end user
  I want to make a call and send a text as SMS.

Scenario: Send Sms
    Given my test setup runs
    When I send SMS "Ask not what your country can do for you. Ask what you can do for your country." from "NumberA" to "NumberB"
    And After waiting for 1 seconds
    Then I should view the SMS "Ask not what your country can do for you. Ask what you can do for your country." from "NumberA" to "NumberB"

Feature: ViewSMS    
  In order to view an SMS 
  As an end user
  I want to list my sms and view the first.

Scenario: View Sms
    Given my test setup runs
    When I send SMS "Actions speak louder than words" from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should view the SMS "Actions speak louder than words" from "NumberB" to "NumberC"

