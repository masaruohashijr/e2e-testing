Feature: MMS    
  In order to my sent MMS messages can be seem by my destination numbers
  And getting a receiving confirmation from them
  As an end user
  I want to send an MMS message

  Scenario: Send an MMS

    Given my test setup runs
      And "NumberB" configured to send MMS "Actions speak louder than words" and media "Avaya.jpg" to "NumberC"
      When I make a call from "NumberA" to "NumberB" 
      Then "NumberB" should be able to view the MMS "I think to myself" with media "Avaya.jpg"