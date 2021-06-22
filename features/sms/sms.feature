Feature: dial    
  In order to make an SMS during a call

Scenario: Sms
    Given my test setup runs
    And "NumberA" configured to send sms
    When I make a call from "NumberA" to "NumberB" 
    Then SMS Status should be sent to call Status URL