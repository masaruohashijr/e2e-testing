Feature: Dial    
  In order to make a seamless call transfer to another number
  As an end user
  I want to make an outgoing dial from an already made and current call

Background: Setup
    Given my test setup runs

Scenario: Dial
    Given "NumberD" configured to dial "NumberBR1"
    When I make a call from "NumberC" to "NumberD"
    Then "NumberBR1" should get the incoming call from "NumberD"