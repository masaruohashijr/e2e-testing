Feature: Release Number    
  In order to release a number from my list
  As an end user
  I want to list numbers and release numbers except "NumberA,NumberB,NumberC,NumberD,NumberE"

  Scenario: List available numbers and view friendly name of one

    Given my test setup runs    
    When I list my numbers
    And I release all my numbers except "NumberA,NumberB,NumberC,NumberD,NumberE"
    Then I should get 5 numbers from my list
