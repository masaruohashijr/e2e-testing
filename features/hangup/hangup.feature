Feature: hangup    
  In order to hangup a call after a certain time in seconds
  As an end user
  I want to have my call automatically off

  Scenario: Hangup a call

    Given my test setup runs
    And "NumberA" configured to hangup after 3 seconds
    When I make a call from "NumberA" to "NumberB"
    Then "NumberA" should get last call duration equals or more than 3