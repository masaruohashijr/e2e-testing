Feature: Gather    
  In order to read text to the Number B (called) using a text-to-speech engine   
  As an end user
  I want that Number A (caller) listen the speech set to be read on Number B.

  Scenario: Gather something said

    Given my test setup runs 
      And "NumberA" configured to say "we shall fight on the beaches" 
      And "NumberB" configured to gather speech 
      When I make a call from "NumberA" to "NumberB" 
      Then "NumberB" should get speech "we shall fight on the beaches" 

Feature: Update_Account
  In order to configure a new friendly name for my account
  As an end user
  I want to view my account's friendly name
  And update a new alias for my account

  Scenario: Update Account

    Given my test setup runs    
    And I update the friendly name for my account to "Masaru"
    When I view my account information
    Then I should get to see "Masaru" as the friendly name for my account


