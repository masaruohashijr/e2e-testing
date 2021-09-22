Feature: ListCalls
As an end user
I want to list my calls
So that I can see the duration of my last call

Scenario: List My Calls

Given my test setup runs
And "NumberC" configured to pause 3 seconds
And append To "NumberC" config hangup
When I make a call from "NumberB" to "NumberC"
And list calls after 3 seconds
Then "NumberB" should get last call duration with more than 2 seconds

