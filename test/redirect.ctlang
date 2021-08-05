Feature: Redirect    
  In order to redirect a call
  As an end user
  I want to make a call to a Number
  And through redirect the call will be off immediately.

  Scenario: Redirect to receive a Speech

    Given my test setup runs 
      And "NumberB" configured to redirect to ping URL
      When I make a call from "NumberA" to "NumberB"
      Then "NumberB" should get a ping request on the URL



