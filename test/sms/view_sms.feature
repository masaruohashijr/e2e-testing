Feature: ViewSMS    
  In order to view an SMS 
  As an end user
  I want to list my sms and view the first.

Scenario: View Sms
    Given my test setup runs
    When I send SMS "Actions speak louder than words" from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should view the SMS "Actions speak louder than words" from "NumberB" to "NumberC"