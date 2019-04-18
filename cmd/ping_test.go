package cmd

import (
	"bytes"
	"errors"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/mriedmann/rocketchat-cli/test"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/url"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func init() {
	Config = viper.New()
	Config.Set("rocketchat.url", "http://localhost:3000")
}

func TestPingCli(t *testing.T) {
	const (
		maxAttempts   = 5
		waitTimeInSec = 1
	)

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		apiController := test.MockedApiController{}
		waitTime := time.Duration(waitTimeInSec) * time.Second
		apiController.On("Ping", maxAttempts, waitTime, true).Return(nil)
		return &apiController
	}
	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs([]string{"ping", "-w", strconv.Itoa(waitTimeInSec), "-m", strconv.Itoa(maxAttempts), "-v"})

	cmd, err := rootCmd.ExecuteC()
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cmd)
	apiController.(*test.MockedApiController).AssertExpectations(t)
}

func TestPingCliFailed(t *testing.T) {
	const (
		maxAttempts   = 5
		waitTimeInSec = 1
	)
	expectedError := errors.New("ping error")

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		apiController := test.MockedApiController{}
		waitTime := time.Duration(waitTimeInSec) * time.Second
		apiController.On("Ping", maxAttempts, waitTime, true).Return(expectedError)
		return &apiController
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs([]string{"ping", "-w", strconv.Itoa(waitTimeInSec), "-m", strconv.Itoa(maxAttempts), "-v"})

	cmd, err := rootCmd.ExecuteC()
	assert.Error(t, err, expectedError)

	output := buf.String()
	assert.Regexp(t, regexp.MustCompile("^Error: "+expectedError.Error()), output)

	assert.NotNil(t, cmd)
	apiController.(*test.MockedApiController).AssertExpectations(t)
}
