Feature: ListUsages    
  In order to list usages
  As an end user
  I want to make a call and get at least one usage.

Scenario: List Usage
    Given my test setup runs
    When I make a call from "NumberA" to "NumberB"
    And After waiting for 2 seconds
    Then I should list at least 1 usage