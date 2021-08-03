Feature: conference
In order to make a conference call
As an end user
I want to dial to a number, create and set a conference room

Scenario: Make a Conference

Given my test setup runs
And "NumberA" configured as conference with size 3
And "NumberB" configured to gather speech
And "NumberC" configured to say "What a wonderful world"
When I make a call from "NumberB" to "NumberA"
And  I make a call from "NumberC" to "NumberA"
Then "NumberB" should get speech "What a wonderful world"


