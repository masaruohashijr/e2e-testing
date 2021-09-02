Feature: View_Account
  In order to configure a new friendly name for my account
  As an end user
  I want to view my account's friendly name
  And update a new alias for my account

  Scenario: View Account

    Given my test setup runs    
    When I view my account information
    Then I should get to see my account balance more than or equal "0"
