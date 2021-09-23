Feature: SendMMS    
  In order to send an MMS during a call
  As an end user
  I want to make a call and send a text as MMS.

Scenario: Send Mms
    Given my test setup runs
    When I send MMS "Actions speak louder than words" from "NumberB" to "NumberC"
    And After waiting for 1 seconds
    Then I should view the MMS "Actions speak louder than words" from "NumberB" to "NumberC"