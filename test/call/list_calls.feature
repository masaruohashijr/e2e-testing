Feature: ListCalls  
  As an end user
  I want to list my calls
  So that I can see the duration of my last call

  Scenario: List My Calls

    Given my test setup runs    
    When  I make a call from "NumberB" to "NumberC"
    And "NumberB" list calls
    Then "NumberB" should get call duration with more than 2 seconds
