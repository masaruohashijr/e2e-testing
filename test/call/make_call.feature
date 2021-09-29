Feature: MakeCall   
  In order to make a call
  As an end user
  I want to make a call from a number to another number

  Scenario: Make a Call

    Given my test setup runs    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 3 seconds
    Then I should get to view a call from "NumberA" to "NumberB" with status "completed"
