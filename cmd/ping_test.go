package cmd

import (
	"bytes"
	"errors"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/mriedmann/rocketchat-cli/test"
	"github.com/stretchr/testify/assert"
	"net/url"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func TestPingCli(t *testing.T) {
	const (
		maxAttempts   = 5
		waitTimeInSec = 1
	)

	c := test.MockedApiController{}

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		waitTime := time.Duration(waitTimeInSec) * time.Second
		c.On("Ping", maxAttempts, waitTime, true).Return(nil)
		return &c
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs([]string{"ping", "-w", strconv.Itoa(waitTimeInSec), "-m", strconv.Itoa(maxAttempts), "-v"})

	cmd, err := rootCmd.ExecuteC()
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cmd)
	c.AssertExpectations(t)
}

func TestPingCliFailed(t *testing.T) {
	const (
		maxAttempts   = 5
		waitTimeInSec = 1
	)
	expectedError := errors.New("ping error")

	c := test.MockedApiController{}

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		waitTime := time.Duration(waitTimeInSec) * time.Second
		c.On("Ping", maxAttempts, waitTime, true).Return(expectedError)
		return &c
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs([]string{"ping", "-w", strconv.Itoa(waitTimeInSec), "-m", strconv.Itoa(maxAttempts), "-v"})

	cmd, err := rootCmd.ExecuteC()
	assert.Error(t, err, expectedError)

	output := buf.String()
	assert.Regexp(t, regexp.MustCompile("^Error: "+expectedError.Error()), output)

	assert.NotNil(t, cmd)
	c.AssertExpectations(t)
}
