Feature: ListCalls  
  As an end user
  I want to make a call and then list my calls
  So that I can list at least 1 call 

  Scenario: List My Calls

    Given my test setup runs    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 3 seconds
    Then I should list at least 1 call


