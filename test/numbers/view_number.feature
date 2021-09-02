Feature: ViewNumber
  As an end user 
  I want to view the details
  So that I can list only the numbers from the list
  
  Scenario: View Number

    Given my test setup runs
      And "NumberA" configured with friendly name as "Zarbat" 
      When I view "NumberA" info
      Then I should get friendly name "Zarbat" on "NumberA"
