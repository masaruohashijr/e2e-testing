Feature: play    
  In order to play a song when someone calls my number
  As an end user
  I want to set a MP3 to play every time another number calls mine.

  Scenario: Play a MP3

    Given my test setup runs
    And "NumberA" configured to play "music.mp3"
    When I make a call from "NumberA" to "NumberB"
    Then "NumberB" should be able to listen 