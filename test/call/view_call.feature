Feature: ViewLastCall

  Scenario: View Last Call

    Given my test setup runs    
    And "NumberB" configured to hang up after 3 seconds    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 3 seconds
    Then I should get last call duration greater than or equal to 3 seconds