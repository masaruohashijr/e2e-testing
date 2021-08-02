Feature: SMS    
  In order to send an SMS during a call
  As an end user
  I want to make a call and send a text as SMS.

Scenario: Sms Message
    Given my test setup runs
    When "NumberA" sends SMS "what we do in life echoes in eternity" to "NumberB"
    Then "NumberB" should be able to view the SMS "what we do in life echoes in eternity"