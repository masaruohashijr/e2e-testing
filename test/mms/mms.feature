Feature: MMS    
  In order to my sent MMS messages can be seem by my destination numbers
  And getting a receiving confirmation from them
  As an end user
  I want to send an MMS message

  Scenario: Send an MMS

    Given my test setup runs
      And created MMS with text "I think to myself" and media "Avaya.jpg"
      When I send MMS from "NumberA" to "NumberB"
      Then "NumberB" should get MMS with text "I think to myself" and media "Avaya.jpg"