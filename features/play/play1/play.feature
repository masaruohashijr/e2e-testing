Feature: play    
  In order to play a song when someone calls my number
  As an end user
  I want to set a tone to play after the call is established

  Scenario: Play a tone

    Given my test setup runs
    And "NumberA" configured to play tone "5000,100,440,800"
    And "NumberB" configured to record calls
    When I make a call from "NumberA" to "NumberB"
    Then "NumberB" should be able to listen to frequencies "440,800"
  
  