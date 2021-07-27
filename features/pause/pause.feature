Feature: pause    
  In order to pause a call for a certain time
  As an end user
  I want to call to a Number and pause some seconds and after hangup.

  Background: setup
    Given my test setup runs

  Scenario: Pause a sequence of sentences
    And "NumberE" configured to pause 3 seconds
    And append To "NumberE" config hangup
    When I make a call from "NumberC" to "NumberE"
    Then "NumberC" should get last call duration more than or equals to 3

