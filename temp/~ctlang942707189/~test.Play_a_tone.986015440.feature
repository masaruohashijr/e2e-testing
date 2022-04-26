Feature: Play a tone
Scenario: Play a tone
Given my test setup runs 
And "NumberA" configured to play tone "5000,10,850"
And "NumberB" configured to record calls for download
When I make a call from "NumberA" to "NumberB"
Then "NumberB" should be able to listen to frequencies "850"
And "NumberA" should be reset
And "NumberB" should be reset
