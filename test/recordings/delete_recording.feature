
Feature: Delete_Recording
  In order to record a call
  As an end user
  I want to try to make a call and record this call until key "#" is pressed.

  Scenario: Delete Recording

    Given my test setup runs
      And "NumberA" configured to pause 3 seconds
      And Append To "NumberA" config say "what we do in life echoes in eternity"
      When I make a call from "NumberA" to "NumberB"
      And I record current call from "NumberA" to "NumberB" for 3 seconds
      And After waiting for 10 seconds
      And I delete all recordings from "NumberA" to "NumberB"
      Then I should list no recording from "NumberA" to "NumberB"