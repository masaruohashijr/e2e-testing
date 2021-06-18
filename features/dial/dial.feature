Feature: dial    
  In order to make a seamless call transfer to another number
  As an end user
  I want to make an outgoing dial from an already made and current call

Scenario: Dial
    Given my test setup runs
    And "NumberA" configured to dial "NumberC"
    When I make a call from "NumberA" to "NumberB"
    Then "NumberC" should get the incoming call from "NumberA"