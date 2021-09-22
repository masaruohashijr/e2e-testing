Feature: SendSMS
In order to send an SMS during a call
As an end user
I want to make a call and send a text as SMS.

Scenario: Sms Message
Given my test setup runs
When I send SMS "Actions speak louder than words" from "NumberA" to "NumberC"
Then "NumberC" should be able to view the SMS "Actions speak louder than words"
