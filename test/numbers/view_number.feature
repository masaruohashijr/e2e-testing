Feature: ViewNumber
  As an end user 
  I want to view the details
  So that I can list only the numbers from the list
  
  Scenario: View Number

    Given my test setup runs
      And "NumberC" configured with VoiceUrl as "http://static.zang.io/ivr/welcome/call.xml" 
      When I view "NumberC" info
      Then I should get VoiceUrl "http://static.zang.io/ivr/welcome/call.xml" on "NumberC"
