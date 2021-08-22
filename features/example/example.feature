Feature: example    
  In order to make a conference call
  As an end user
  I want to dial to a number, create and set a conference room

  Scenario: Make an Example

Given my test setup runs #-> myTestSetupRuns()
  And "NumberA" configured to say "what we do in life echoes in eternity" 
  And "NumberB" configured to gather speech 
  When I make a call from "NumberA" to "NumberB" 
  Then "NumberB" should get speech "what we do in life echoes in eternity" 