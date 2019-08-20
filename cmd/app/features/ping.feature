Feature: Ping the scorekeeper microservice
    As a user of the scorekeeper microservice
    I want to have a simple check that the microservice is running
    So that I know it is up

    Scenario: Ping the scorekeeper microservice
        Given that the scorekeeper service is running
        Then I can ping the scorekeeper service
