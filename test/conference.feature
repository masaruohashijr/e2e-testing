Feature: conference    
  In order to make a conference call
  As an end user
  I want to dial to a number, create and set a conference room
  
  Scenario: Make a Conference

    Given my test setup runs
      And "NumberC" configured as conference "Zarbat" with size 3
      When I make a call from "NumberB" to "NumberC"
      And  I make a call from "NumberA" to "NumberC"
      Then "NumberB" should be enter conference "Zarbat" 

