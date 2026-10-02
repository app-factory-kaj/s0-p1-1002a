Feature: Greeting

  @story-1
  Rule: The service returns a personalized JSON greeting for a given name

    Scenario: Greeting a named caller
      Given the greeter service is running
      When an API Consumer sends a GET request to /hello with name "Ada"
      Then the response is a JSON greeting for "Ada"

  @story-2
  Rule: The service returns a default greeting when no name is given

    Scenario: Omitting the name parameter
      Given the greeter service is running
      When an API Consumer sends a GET request to /hello with no name parameter
      Then the response is a JSON greeting for "World"

    Scenario: Supplying an empty name parameter
      Given the greeter service is running
      When an API Consumer sends a GET request to /hello with name ""
      Then the response is a JSON greeting for "World"
