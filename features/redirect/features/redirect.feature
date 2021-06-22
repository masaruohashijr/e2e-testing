Feature: redirect    
  In order to redirect a call
  As an end user
  I want to make a call to a Number and listen to a message from the redirected Url.

  Scenario: Redirect to receive a Speech

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberA" configured to gather speech #-> configuredToGatherSpeech(numberB string, text string)
      And "NumberB" configured to say "You will be redirected." #-> configuredToSay(numberB string, text string)
      And append to "NumberB" config redirect to "redirect_url" #-> appendToConfigRedirectTo(numberB string, redirectTo string)
      And "redirect_url" configures "NumberC" to say "You was redirected." #-> configuresToSay(url string, numerC string, text string)
      When I make a call from "NumberA" to "NumbeB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberA" should get speech "You will be redirected. You was redirected." #-> shouldGetSpeech(numberA string, speech string)