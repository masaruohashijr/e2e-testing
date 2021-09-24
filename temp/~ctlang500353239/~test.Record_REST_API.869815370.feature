Feature: Record_REST_API
In order to record a call
As an end user
I want to try to make a call and record this call until key "#" is pressed.

Scenario: Record a call

Given my test setup runs
And "NumberA" configured to say "what we do in life echoes in eternity"
When I make a call from "NumberA" to "NumberB"
And I record current call from "NumberA" to "NumberB" for 5 seconds
Then I should list at least 1 recording from "NumberB" to "NumberC"

