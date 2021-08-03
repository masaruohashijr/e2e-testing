Feature: Purchase a Number
In order to list available numbers from one area code
As an end user
I want to list at least one number from this area

Scenario: List available numbers and buy one

Given my test setup runs
When I list all available numbers
Then I should get to buy 1 from list


