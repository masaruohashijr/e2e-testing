Feature: List Number    
  In order to list my numbers
  As an end user
  I want to see a list of my numbers

  Scenario: List available numbers and view friendly name of one

    Given my test setup runs
    When I list my numbers
    Then I should get 5 numbers
