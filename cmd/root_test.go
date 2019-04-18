package cmd

import (
	"bytes"
	"github.com/mitchellh/go-homedir"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/mriedmann/rocketchat-cli/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/url"
	"os"
	"strings"
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

func runCmd(args []string) (*bytes.Buffer, *cobra.Command, error) {
	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs(args)
	cmd, err := rootCmd.ExecuteC()
	return buf, cmd, err
}

func setMockedConfigDefault(config *test.MockedConfig) {
	config.On("SetConfigType", mock.AnythingOfType("string"))
	config.On("SetEnvPrefix", mock.AnythingOfType("string"))

	config.On("SetConfigFile", cfgFile).Maybe()
	config.On("AddConfigPath", mock.AnythingOfType("string")).Maybe()
	config.On("SetConfigName", mock.AnythingOfType("string")).Maybe()

	config.On("SetEnvKeyReplacer", mock.AnythingOfType("*strings.Replacer"))
	config.On("AutomaticEnv")
	config.On("ReadInConfig").Return(nil)

	config.On("IsSet", "rocketchat.url").Return(true)
	config.On("GetString", "rocketchat.url").Return(rocketchatUrlRaw)

	config.On("IsSet", mock.AnythingOfType("string")).Return(false)
	config.On("GetString", mock.AnythingOfType("string")).Return("")
}

func TestMain(m *testing.M) {
	config := test.MockedConfig{}
	setMockedConfigDefault(&config)
	Config = &config

	retCode := m.Run()
	Config = nil
	os.Exit(retCode)
}

var mockedConfig test.MockedConfig

func TestRootInitConfig(t *testing.T) {
	home, _ := homedir.Dir()

	mockedConfig = test.MockedConfig{}

	mockedConfig.On("SetConfigType", "yaml")
	mockedConfig.On("SetEnvPrefix", "rccli")
	mockedConfig.On("AddConfigPath", home)
	mockedConfig.On("AddConfigPath", ".")
	mockedConfig.On("SetConfigName", ".rocketchat-cli")

	mockedConfig.On("SetEnvKeyReplacer", strings.NewReplacer(".", "_"))
	mockedConfig.On("AutomaticEnv")
	mockedConfig.On("ReadInConfig").Return(nil)

	mockedConfig.On("IsSet", "rocketchat.url").Return(true)
	mockedConfig.On("GetString", "rocketchat.url").Return(rocketchatUrlRaw)

	mockedConfig.On("GetString", "user.email").Return(testCredentials.Email)
	mockedConfig.On("GetString", "user.id").Return(testCredentials.ID)
	mockedConfig.On("GetString", "user.token").Return(testCredentials.Token)
	mockedConfig.On("GetString", "user.password").Return(testCredentials.Password)

	Config = &mockedConfig

	initConfig()

	mockedConfig.AssertExpectations(t)
}

func TestRootInitConfigFile(t *testing.T) {
	cfgFile = "testconfigfile"

	config := test.MockedConfig{}
	config.On("SetConfigType", "yaml")
	config.On("SetEnvPrefix", "rccli")

	config.On("SetConfigFile", cfgFile)

	config.On("SetEnvKeyReplacer", strings.NewReplacer(".", "_"))
	config.On("AutomaticEnv")
	config.On("ReadInConfig").Return(nil)

	config.On("IsSet", "rocketchat.url").Return(true)
	config.On("GetString", "user.email").Return(testCredentials.Email)
	config.On("GetString", "user.id").Return(testCredentials.ID)
	config.On("GetString", "user.token").Return(testCredentials.Token)
	config.On("GetString", "user.password").Return(testCredentials.Password)
	config.On("GetString", "rocketchat.url").Return(rocketchatUrlRaw)

	Config = &config

	initConfig()

	config.AssertExpectations(t)
}

func TestRootInitConfigNoUrl(t *testing.T) {
	config := test.MockedConfig{}
	config.On("SetConfigType", mock.AnythingOfType("string"))
	config.On("SetEnvPrefix", mock.AnythingOfType("string"))
	config.On("SetConfigFile", cfgFile).Maybe()
	config.On("AddConfigPath", mock.AnythingOfType("string")).Maybe()
	config.On("SetConfigName", mock.AnythingOfType("string")).Maybe()

	config.On("SetEnvKeyReplacer", mock.AnythingOfType("*strings.Replacer"))
	config.On("AutomaticEnv")
	config.On("ReadInConfig").Return(nil)

	config.On("IsSet", "rocketchat.url").Return(false)
	Config = &config

	assert.PanicsWithValue(t, "config error - rocketchat.url not set", initConfig)
}

func TestRootInitConfigInvalidUrl(t *testing.T) {
	var invalidUrl = ":://invalidurl"

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		t.Fail()
		return nil
	}

	config := test.MockedConfig{}
	config.On("SetConfigType", mock.AnythingOfType("string"))
	config.On("SetEnvPrefix", mock.AnythingOfType("string"))
	config.On("SetConfigFile", cfgFile).Maybe()
	config.On("AddConfigPath", mock.AnythingOfType("string")).Maybe()
	config.On("SetConfigName", mock.AnythingOfType("string")).Maybe()

	config.On("SetEnvKeyReplacer", mock.AnythingOfType("*strings.Replacer"))
	config.On("AutomaticEnv")
	config.On("ReadInConfig").Return(nil)

	config.On("IsSet", "rocketchat.url").Return(true)
	config.On("GetString", "rocketchat.url").Return(invalidUrl)

	config.On("GetString", mock.AnythingOfType("string")).Return("")
	Config = &config

	assert.Panics(t, initConfig)
}
