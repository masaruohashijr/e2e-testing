Feature: ViewLastCall

  Scenario: View Last Call

    Given my test setup runs    
    And "NumberC" configured to pause 2 seconds
    And Append To "NumberC" config hang up
    When I make a call from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should get last call duration greater than or equal to 2 seconds