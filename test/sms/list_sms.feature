Feature: ListSMS    
  In order to send an SMS during a call
  As an end user
  I want to make a call and send a text as SMS.

Scenario: List Sms
    Given my test setup runs
    When I send SMS "Actions speak louder than words" from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should list at least 1 SMS from "NumberB" to "NumberC"