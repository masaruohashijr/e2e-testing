Feature: sms    
  In order to send an SMS during a call
  As an end user
  I want to make a call and send a text as SMS.

Scenario: Sms Message
    Given my test setup runs
    When "NumberA" sends SMS "what we do in life echoes in eternity" to "NumberB"
    Then "NumberB" should be able to view the SMS "what we do in life echoes in eternity"

#Scenario: Sms Status
#    Given my test setup runs
#    And "NumberA" configured to send sms to "NumberB"
#    When I make a call from "NumberA" to "NumberB" 
#    Then SMS Status should be sent to call Status URL