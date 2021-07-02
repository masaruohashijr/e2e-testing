Feature: pause    
  In order to pause a call for a certain time
  As an end user
  I want to call to a Number and pause some seconds and after hangup.

  Scenario: Pause a sequence of sentences

    Given my test setup runs
    And "NumberB" configured to pause 3 seconds
    And append To "NumberB" config hangup
    When I make a call from "NumberA" to "NumberB"
    Then "NumberA" should get last call duration more than or equals to 3
