Feature: MakeCall   
  In order to make a call
  As an end user
  I want to make a call from a number to another number

  Scenario: Make a Call

    Given my test setup runs
    And "NumberB" configured to hang up after 2 seconds    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 10 seconds
    Then I should get to view a call from "NumberA" to "NumberB" with status "completed"

Feature: ListCalls  
  As an end user
  I want to make a call and then list my calls
  So that I can list at least 1 call 

  Scenario: List My Calls

    Given my test setup runs    
    And "NumberB" configured to hang up after 2 seconds    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 3 seconds
    Then I should list at least 1 call

Feature: ViewLastCall

  Scenario: View Last Call

    Given my test setup runs    
    And "NumberB" configured to hang up after 3 seconds    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 3 seconds
    Then I should get last call duration greater than or equal to 3 seconds

