Feature: play    
  In order to play a song when someone calls my number
  As an end user
  I want to set a MP3 to play every time another number calls mine.

  Scenario: Play a MP3

    Given my test setup runs
    And "NumberF" configured to play "sample3"
    And "NumberE" configured to gather speech
    When I make a call from "NumberF" to "NumberE"
    Then "NumberE" should be able to listen 
  
  