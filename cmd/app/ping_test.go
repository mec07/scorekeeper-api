package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/mec07/rununtil"
	"github.com/pkg/errors"
)

// This is just to run godog when running go test
func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "progress",
		Paths:     []string{"features"},
		Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func thatTheScorekeeperServiceIsRunning() error {
	return nil
}

func iCanPingTheScorekeeperService() error {
	url := "http://localhost:8080/ping"
	res, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "pinging %s", url)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "reading response body")
	}
	if string(body) != "pong" {
		return fmt.Errorf(`expected response to be "pong", got: "%s"`, string(body))
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	// Before & After steps
	var killServer context.CancelFunc
	s.BeforeScenario(func(interface{}) {
		killServer = rununtil.Killed(main)
	})

	s.AfterScenario(func(interface{}, error) {
		killServer()
	})

	// Given steps
	s.Step(`^that the scorekeeper service is running$`, thatTheScorekeeperServiceIsRunning)

	// When steps

	// Then steps
	s.Step(`^I can ping the scorekeeper service$`, iCanPingTheScorekeeperService)
}
