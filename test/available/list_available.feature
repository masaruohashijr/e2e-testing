Feature: List_Available_Numbers
  In order to list at least one available number 
  As an end user
  I want to list all available numbers

  Scenario: List available numbers 

    Given my test setup runs    
    When I list all available numbers
    Then I should get to see more than or equal to 1 from list 
