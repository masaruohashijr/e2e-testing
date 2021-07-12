Feature: pause    
  In order to pause a call for a certain time
  As an end user
  I want to call to a Number and pause some seconds and after hangup.

  Background: setup
    Given my test setup runs

  Scenario: Pause a sequence of sentences
    And "NumberA" configured to pause 3 seconds
    And append To "NumberA" config hangup
    When I make a call from "NumberB" to "NumberA"
    Then "NumberB" should get last call duration more than or equals to 3

