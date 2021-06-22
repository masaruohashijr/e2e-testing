Feature: number    
  In order to try a sequence of numbers until connect
  As an end user
  I want to call to a Number
  And this current call, in case the previous number does not connect, trys to connect to a next number in sequence.

  Scenario: Call multiple numbers in a sequence

    Given my test setup runs #-> myTestSetupRuns()
    And "NumberB" configured to dial "NumberC" #-> configuredToDial(numberB string, numberC string)
    And append to "NumberB" config dial "NumberD" #-> appendToConfigDial(numberB string, numberD string)
    And "NumberC" configured to reject with "not-in-service" reason #-> configuredToRejectWith(numberC string, reason string)
    When I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberB string, numberA string)
    Then "NumberD" should get the incoming call from "NumberA" #-> shouldGetTheIncomingCallFrom(numberA string, numberB string)
