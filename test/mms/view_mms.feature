Feature: ViewMMS    
  In order to view an MMS 
  As an end user
  I want to list my sms and view the first.

Scenario: View Sms
    Given my test setup runs
    When I send MMS "Actions speak louder than words" from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should view the MMS "Actions speak louder than words" from "NumberB" to "NumberC"