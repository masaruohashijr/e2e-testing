Feature: SMS
In order to send an SMS during a call
As an end user
I want to make a call and send a text as SMS.

Scenario: Sms Message
Given my test setup runs
And "NumberB" configured to send SMS "Actions speak louder than words" to "NumberC"
When I make a call from "NumberA" to "NumberB"
Then "NumberC" should be able to view the SMS "Actions speak louder than words"


#Feature: Purchase a Number
#  In order to list available numbers from one area code
#  As an end user
#  I want to list at least one number from this area

#Scenario: List available numbers and buy one

#    Given my test setup runs
#    When I list all available numbers
#    Then I should get to buy 1 from list

