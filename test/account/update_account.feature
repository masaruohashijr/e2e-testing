Feature: Update_Account
  In order to configure a new friendly name for my account
  As an end user
  I want to view my account's friendly name
  And update a new alias for my account

  Scenario: Update Account

    Given my test setup runs    
    And I update the friendly name for my account to "Zarbat"
    When I view my account information
    Then I should get to see "Zarbat" as the friendly name for my account
