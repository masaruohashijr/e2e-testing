Feature: ping    
  In order to send data from a current call to a Webhook
  As an end user
  I want to call to a Number and ping a URL.

  Scenario: Ping a URL
    Given my test setup runs
    And "NumberA" configured to ping URL
    When I make a call from "NumberA" to "NumberB"
    Then should get a ping request on the URL