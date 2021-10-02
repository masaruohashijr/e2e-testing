Feature: ViewUsage
  In order to view usage
  As an end user
  I want to make a call and view the total cost more than zero.

Scenario: List Usage
    Given my test setup runs
    And "NumberB" configured to hang up after 2 seconds    
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 4 seconds
    Then I should view the total cost usage more than 0