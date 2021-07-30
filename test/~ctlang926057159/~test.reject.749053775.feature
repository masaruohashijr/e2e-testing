Feature: reject
In order to reject a call
As an end user
I want to try to make a call to a Number and other number rejects.

Scenario: Call Reject
Given my test setup runs
And "NumberB" configured to reject call
When I make a call from "NumberA" to "NumberB"
Then "NumberA" should get call cancel status



