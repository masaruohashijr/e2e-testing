Feature: conference    
  In order to make a conference call
  As an end user
  I want to dial to a number, create and set a conference room
  
  Scenario: Make a Conference

    Given my test setup runs
      And "NumberC" configured as conference "Zarbat" with size 3
      When I make a call from "NumberB" to "NumberC"
      And  I make a call from "NumberA" to "NumberC"
      Then "NumberB" should had entered conference "Zarbat" 


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


Feature: Hangup    
  In order to hang up a call after a certain time in seconds
  As an end user
  I want to have my call automatically off

Background: Setup
    Given my test setup runs

  Scenario: Hangup a call
    And "NumberB" configured to hang up after 3 seconds
    When I make a call from "NumberA" to "NumberB"
    Then "NumberB" should get last call duration more than or equals to 3


Feature: Number    
  In order to make a seamless call transfer to another number
  As an end user
  I want to make an outgoing dial from an already made and current call
  And send digits to the dialed number

Background: Setup
    Given my test setup runs

Scenario: Number
    Given "NumberB" configured to dial and send digits "wwww1234#" to "NumberC"
    And "NumberC" configured to gather digits until "#"
    When I make a call from "NumberA" to "NumberB"
    Then "NumberC" should get digits "1234" from "NumberB"


Feature: Pause    
  In order to pause a call for a certain time
  As an end user
  I want to call to a Number and pause some seconds and after hang up.

  Background: setup
    Given my test setup runs

  Scenario: Pause a sequence of sentences
    And "NumberD" configured to pause 3 seconds
    And Append To "NumberD" config hang up
    When I make a call from "NumberC" to "NumberD"
    Then "NumberD" should get last call duration more than or equals to 3




Feature: Ping    
  In order to send data from a current call to a Webhook
  As an end user
  I want to call to a Number and ping a URL.

  Background: setup
    Given my test setup runs

  Scenario: Ping a URL
    And "NumberB" configured to ping URL
    When I make a call from "NumberA" to "NumberB"
    Then "NumberB" should get a ping request on the URL


Feature: Play    
  In order to play a tone with frequencies
  As an end user
  I want to set a tone to play after the call is established
  And should be able to record and extract these frequencies

  Background: setup
    Given my test setup runs 
    
  Scenario: Play a tone
    And "NumberD" configured to play tone "5000,10,850"
    And "NumberE" configured to record calls for download
    When I make a call from "NumberD" to "NumberE"
    Then "NumberE" should be able to listen to frequencies "850"
    And "NumberD" should be reset
    And "NumberE" should be reset
  


Feature: PlayLastRecording    
  In order to listen the last recording of my calls
  As an end user
  I want to call a number and set to play the last recording.

  Scenario: Record Call

    Given my test setup runs
      And "NumberA" configured to say "This is the last recording"
      And "NumberB" configured to record calls
      When I make a call from "NumberA" to "NumberB"
      Then "NumberB" should get speech "This is the last recording"

  Scenario: Play Last Recording

    Given my test setup runs
      And "NumberA" configured to gather speech 
      And "NumberB" configured to play last recording
      And I make a call from "NumberA" to "NumberB"
      Then "NumberA" should get speech "This is the last recording"
      And "NumberA" should be reset
      And "NumberB" should be reset



Feature: Record    
  In order to record a call
  As an end user
  I want to try to make a call and record this call until key "#" is pressed.

  Scenario: Pause a sequence of sentences

    Given my test setup runs
      And "NumberA" configured to say "what we do in life echoes in eternity"
      And "NumberB" configured to record calls
      When I make a call from "NumberA" to "NumberB"
      Then "NumberB" should get transcription "what we do in life echoes in eternity"      


Feature: Redirect    
  In order to redirect a call
  As an end user
  I want to make a call to a Number
  And through redirect the call will be off immediately.

  Scenario: Redirect to receive a Speech

    Given my test setup runs 
      And "NumberB" configured to redirect to ping URL
      When I make a call from "NumberA" to "NumberB"
      Then "NumberB" should get a ping request on the URL



Feature: Reject    
  In order to reject a call
  As an end user
  I want to try to make a call to a Number and other number rejects.

  Scenario: Call Reject 
    Given my test setup runs
    And "NumberB" configured to reject call
    When I make a call from "NumberA" to "NumberB"
    Then "NumberB" should get call cancel status



Feature: Say    
  In order to read text to the Number B (called) using a text-to-speech engine   
  As an end user
  I want that Number A (caller) listen the speech set to be read on Number B.

  Scenario: Say something

    Given my test setup runs 
      And "NumberA" configured to say "what we do in life echoes in eternity" 
      And "NumberB" configured to gather speech 
      When I make a call from "NumberA" to "NumberB" 
      Then "NumberB" should get speech "what we do in life echoes in eternity" 



Feature: SMS    
  In order to send an SMS during a call
  As an end user
  I want to make a call and send a text as SMS.

Scenario: Sms Message
    Given my test setup runs
    And "NumberB" configured to send SMS "Actions speak louder than words" to "NumberC"
    When I make a call from "NumberA" to "NumberB" 
    Then "NumberC" should be able to view the SMS "Actions speak louder than words"

Feature: Dial    
  In order to make a seamless call transfer to another number
  As an end user
  I want to make an outgoing dial from an already made and current call

Background: Setup
    Given my test setup runs

Scenario: Dial
    Given "NumberD" configured to dial "NumberBR1"
    When I make a call from "NumberC" to "NumberD"
    Then "NumberBR1" should get the incoming call from "NumberD"

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


Feature: Purchase a Number    
  In order to list available numbers from one area code
  As an end user
  I want to list at least one number from this area

Scenario: List available numbers and buy one

    Given my test setup runs    
    When I list all available numbers
    Then I should get to buy 1 from list 
