
Feature: List_Notification
  In order to list notifications
  As an end user
  I want to list the notifications after making a call.

  Scenario: List Notifications

    Given my test setup runs
      And "NumberA" configured to pause 3 seconds
      And Append To "NumberA" config say "what we do in life echoes in eternity"
      When I make a call from "NumberA" to "NumberB"
      Then I should list at least 1 notification from "NumberA" to "NumberB"