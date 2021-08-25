Feature: WriteMyName
  In order to write my name
  As an end user
  I want to see my name

  Scenario: Write My Name

    Given my test setup runs
    When I want to write my name "Masaru"
    Then I should see "Aron" on console
