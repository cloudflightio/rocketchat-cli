package cmd

import (
	"bytes"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/mriedmann/rocketchat-cli/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/url"
	"testing"
)

var testCredentials = models.UserCredentials{
	ID:       "",
	Token:    "",
	Email:    "admin@test.com",
	Password: "testpassword",
}

var (
	rocketchatUrlRaw = "http://localhost:3000"
)

func NewMockedConfigController(_ string, _ bool) controllers.ConfigController {
	c := test.MockedConfigController{}
	c.On("IsSet", "rocketchat.url").Return(true)
	c.On("GetString", "rocketchat.url").Return(rocketchatUrlRaw)

	c.On("GetString", "user.email").Return(testCredentials.Email)
	c.On("GetString", "user.id").Return(testCredentials.ID)
	c.On("GetString", "user.token").Return(testCredentials.Token)
	c.On("GetString", "user.password").Return(testCredentials.Password)
	return &c
}

func runCmd(args []string) (*bytes.Buffer, *cobra.Command, error) {
	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs(args)
	cmd, err := rootCmd.ExecuteC()
	return buf, cmd, err
}

func TestRootInitConfig(t *testing.T) {
	c := test.MockedConfigController{}

	ConfigControllerFactory = NewMockedConfigController

	initConfig()

	c.AssertExpectations(t)
}

func TestRootInitConfigFile(t *testing.T) {
	cfgFile = "testconfigfile"

	c := test.MockedConfigController{}

	ConfigControllerFactory = NewMockedConfigController

	initConfig()

	c.AssertExpectations(t)
}

func TestRootInitConfigNoUrl(t *testing.T) {
	config := test.MockedConfigController{}

	ConfigControllerFactory = func(s string, b bool) controllers.ConfigController {
		config.On("IsSet", "rocketchat.url").Return(false)
		return &config
	}

	assert.PanicsWithValue(t, "config error - rocketchat.url not set", initConfig)
}

func TestRootInitConfigInvalidUrl(t *testing.T) {
	var invalidUrl = ":://invalidurl"

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		t.Fail()
		return nil
	}

	c := test.MockedConfigController{}
	ConfigControllerFactory = func(s string, b bool) controllers.ConfigController {

		c.On("IsSet", "rocketchat.url").Return(true)
		c.On("GetString", "rocketchat.url").Return(invalidUrl)

		c.On("GetString", mock.AnythingOfType("string")).Return("")
		return &c
	}

	assert.Panics(t, initConfig)
}

func TestRootInitConfigBadUrl(t *testing.T) {
	var invalidUrl = "invalidurl"

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		t.Fail()
		return nil
	}

	c := test.MockedConfigController{}
	ConfigControllerFactory = func(s string, b bool) controllers.ConfigController {

		c.On("IsSet", "rocketchat.url").Return(true)
		c.On("GetString", "rocketchat.url").Return(invalidUrl)

		c.On("GetString", mock.AnythingOfType("string")).Return("")
		return &c
	}

	assert.Panics(t, initConfig)
}
