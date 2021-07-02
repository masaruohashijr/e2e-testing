Feature: play    
  In order to play a tone with frequencies
  As an end user
  I want to set a tone to play after the call is established
  And should be able to record and extract these frequencies

  Scenario: Play a tone

    Given my test setup runs
    And "NumberA" configured to play tone "5000,10,1050"
    And "NumberB" configured to record calls
    When I make a call from "NumberA" to "NumberB"
    Then "NumberB" should be able to listen to frequencies "1050"
  
  