Feature: Ping the scorekeeper microservice
    In order to know that the microservice is running
    As a user of the scorekeeper microservice
    I need to have a simple check that the microservice is running

    Scenario: Ping the scorekeeper microservice
        Given that the scorekeeper service is running
        Then I can ping the scorekeeper service
