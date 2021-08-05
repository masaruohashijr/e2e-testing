Feature: Dial
In order to make a seamless call transfer to another number
As an end user
I want to make an outgoing dial from an already made and current call

Background: Setup
Given my test setup runs

Scenario: Dial
Given "NumberB" configured to dial "NumberC"
And Append To "NumberB" config hang up
When I make a call from "NumberA" to "NumberB"
Then "NumberC" should get the incoming call from "NumberB"
