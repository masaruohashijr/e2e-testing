Feature: pause    
  In order to pause a sequence of phrases
  As an end user
  I want to call to a Number and say some words or phrases pausing some seconds between each sentence.

  Scenario: Pause a sequence of sentences

    Given my test setup runs
    And "NumberB" configured to pause 5 seconds
    And append To "NumberB" config hangup
    When I make a call from "NumberA" to "NumberB"
    Then "NumberA" should get last call duration equals to 5
