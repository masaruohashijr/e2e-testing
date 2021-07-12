Feature: ping    
  In order to send data from a current call to a Webhook
  As an end user
  I want to call to a Number and ping a URL.

  Background: setup
    Given my test setup runs

  Scenario: Ping a URL
    And "NumberB" configured to ping URL
    When I make a call from "NumberA" to "NumberB"
    Then should get a ping request on the URL