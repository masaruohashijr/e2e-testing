Feature: Pause
In order to pause a call for a certain time
As an end user
I want to call to a Number and pause some seconds and after hang up.

Background: setup
Given my test setup runs

Scenario: Pause a sequence of sentences
And "NumberD" configured to pause 3 seconds
And Append To "NumberD" config hang up
When I make a call from "NumberC" to "NumberD"
Then "NumberD" should get last call duration more than or equals to 3




