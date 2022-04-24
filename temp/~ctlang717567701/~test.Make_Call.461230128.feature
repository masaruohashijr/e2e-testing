Feature: Make Call
Scenario: Make Call
 Given my test setup runs
 And "NumberB" configured to hang up after 2 seconds
 When I make a call from "NumberA" to "NumberB"
 And After waiting for 10 seconds
 Then I should get to view a call from "NumberA" to "NumberB" with status "completed"

