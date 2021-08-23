Feature: Release Numbers
  As an end user 
  I want to release all my numbers except a list 
  So that I can list only the numbers from the list
  
  Scenario: Release Numbers

    Given my test setup runs
      When I release all my numbers except "NumberA,NumberB,NumberC,NumberD,NumberE"
      Then I should list my numbers as "NumberA,NumberB,NumberC,NumberD,NumberE"
