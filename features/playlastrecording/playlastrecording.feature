Feature: play last recording    
  In order to listen the last recording of my calls
  As an end user
  I want to call a number and set to play the last recording.

  Scenario: Play Last Recording

    Given my test setup runs
      And "NumberA" configured to say "This is the last recording"
      And "NumberB" configured to record calls
      When I make a call from "NumberA" to "NumbeB"
      And "NumberB" configured to play last recording
      Then "NumberA" should get transcription "This is the last recording"