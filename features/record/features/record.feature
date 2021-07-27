Feature: record    
  In order to record a call
  As an end user
  I want to try to make a call and record this call until key "#" is pressed.

  Scenario: Pause a sequence of sentences

    Given my test setup runs
      And "NumberA" configured to say "what we do in life echoes in eternity"
      And "NumberB" configured to record calls
      When I make a call from "NumberA" to "NumberB"
      Then should get transcription "what we do in life echoes in eternity"