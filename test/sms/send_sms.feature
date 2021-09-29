Feature: SendSMS    
  In order to send an SMS during a call
  As an end user
  I want to make a call and send a text as SMS.

Scenario: Send Sms
    Given my test setup runs
    When I send SMS "Ask not what your country can do for you. Ask what you can do for your country." from "NumberA" to "NumberB"
    And After waiting for 1 seconds
    Then I should view the SMS "Ask not what your country can do for you. Ask what you can do for your country." from "NumberA" to "NumberB"