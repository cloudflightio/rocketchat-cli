package cmd

import (
	"errors"
	"github.com/cloudflightio/rocketchat-cli/cmd"
	"github.com/cloudflightio/rocketchat-cli/controllers"
	"github.com/cloudflightio/rocketchat-cli/models"
	"github.com/cloudflightio/rocketchat-cli/test"
	"github.com/stretchr/testify/assert"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

var vm = models.CreateUserViewModel{
	Name:     "Test Zerozeroone",
	Email:    "test001@catalysts.cc",
	Password: "test",
	Username: "test001",
}

func TestCreateUserCli(t *testing.T) {
	cmd.ConfigControllerFactory = NewMockedConfigController
	apiController := test.MockedApiController{}
	cmd.ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		apiController.On("CreateUser", &vm).Return(nil)
		return &apiController
	}

	args := []string{"createUser", "-n", vm.Name, "-u", vm.Username, "-e", vm.Email, "-p", vm.Password}

	buf, cli, err := runCmd(args)
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cli)
	apiController.AssertExpectations(t)
}

func TestCreateUserCliCreateFailed(t *testing.T) {
	expectedError := errors.New("create error")

	cmd.ConfigControllerFactory = NewMockedConfigController
	apiController := test.MockedApiController{}
	cmd.ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		apiController.On("CreateUser", &vm).Return(expectedError)
		return &apiController
	}

	args := []string{"createUser", "-n", vm.Name, "-u", vm.Username, "-e", vm.Email, "-p", vm.Password}

	buf, cli, err := runCmd(args)

	assert.Error(t, err, expectedError)

	output := buf.String()
	assert.Regexp(t, regexp.MustCompile("^Error: "+expectedError.Error()), output)

	assert.NotNil(t, cli)
	apiController.AssertExpectations(t)
}

func TestCreateUserCliIgnoreExisting(t *testing.T) {
	vm := models.CreateUserViewModel(vm)
	vm.IgnoreExisting = true

	cmd.ConfigControllerFactory = NewMockedConfigController
	apiController := test.MockedApiController{}
	cmd.ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		apiController.On("CreateUser", &vm).Return(nil)
		return &apiController
	}

	args := []string{"createUser", "-n", vm.Name, "-u", vm.Username, "-e", vm.Email, "-p", vm.Password, "-i"}

	buf, cli, err := runCmd(args)
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cli)
	apiController.AssertExpectations(t)
}

func TestCreateUserCliRoles(t *testing.T) {
	cmd.IgnoreExisting = false //TODO: find solution for global var "bleeding"
	vm := models.CreateUserViewModel(vm)
	vm.Roles = []string{"bot", "admin"}

	cmd.ConfigControllerFactory = NewMockedConfigController
	apiController := test.MockedApiController{}
	cmd.ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		apiController.On("CreateUser", &vm).Return(nil)
		return &apiController
	}

	args := []string{"createUser", "-n", vm.Name, "-u", vm.Username, "-e", vm.Email, "-p", vm.Password, "-r", strings.Join(vm.Roles, ",")}

	buf, cli, err := runCmd(args)
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cli)
	apiController.AssertExpectations(t)
}
