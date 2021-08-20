Feature: PlayLastRecording    
  In order to listen the last recording of my calls
  As an end user
  I want to call a number and set to play the last recording.

  Scenario: Record Call

    Given my test setup runs
      And "NumberA" configured to say "This is the last recording"
      And "NumberB" configured to record calls
      When I make a call from "NumberA" to "NumberB"
      Then "NumberB" should get speech "This is the last recording"

  Scenario: Play Last Recording

    Given my test setup runs
      And "NumberA" configured to gather speech 
      And "NumberB" configured to play last recording
      And I make a call from "NumberA" to "NumberB"
      Then "NumberA" should get speech "This is the last recording"
      And "NumberA" should be reset
      And "NumberB" should be reset
