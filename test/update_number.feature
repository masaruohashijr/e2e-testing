Feature: Update a Number    
  In order to view the friendly name of a number
  As an end user
  I want to list numbers and view the friendly name the first one

  Scenario: Update a friendly name number and view friendly name

    Given my test setup runs    
    When I list my numbers
    And "NumberA" configured with friendly name as "Zarbat"
    Then I should get friendly name "Zarbat" on "NumberA"
