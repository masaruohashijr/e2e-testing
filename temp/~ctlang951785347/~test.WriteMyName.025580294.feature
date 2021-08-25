Feature: WriteMyName
In order to write my name
As an end user
I want to see my name

Scenario: List available numbers and view friendly name of one

Given my test setup runs
When I want to write my name "Aron"
Then I should see "Aron" on console

