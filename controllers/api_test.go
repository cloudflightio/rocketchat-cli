package controllers

import (
	"errors"
	"fmt"
	sdkModels "github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/mriedmann/rocketchat-cli/test"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

var credentials = models.UserCredentials{Email: "admin@catalysts.cc", Password: "admin"}

var request = sdkModels.CreateUserRequest{
	Name:         "test001",
	Email:        "test001@catalysts.cc",
	Password:     "test",
	Username:     "test001",
	CustomFields: nil,
}

var createUserVM = models.CreateUserViewModel{
	Name:     request.Name,
	Email:    request.Email,
	Password: request.Password,
	Username: request.Username,
}

func getTestableApiController() (*SdkApiController, *test.MockedClient, *sdkModels.UserCredentials) {
	uc := sdkModels.UserCredentials{Email: credentials.Email, Password: credentials.Password}
	client := test.MockedClient{}
	c := SdkApiController{
		client:      &client,
		credentials: &uc,
	}
	return &c, &client, &uc
}

func TestNewSdkApiController(t *testing.T) {
	serverUrl, _ := url.Parse("http://localhost:3000")

	expectedClient := rest.Client{
		Debug:    true,
		Host:     serverUrl.Hostname(),
		Protocol: serverUrl.Scheme,
		Port:     serverUrl.Port(),
		Version:  "v1",
	}
	uc := sdkModels.UserCredentials{Email: credentials.Email, Password: credentials.Password}

	c := NewSdkApiController(serverUrl, true, &credentials).(*SdkApiController)

	assert.Equal(t, &expectedClient, c.client)
	assert.Equal(t, &uc, c.credentials)
}

func TestCreateUser(t *testing.T) {
	c, client, uc := getTestableApiController()

	client.On("Login", uc).Return(nil)
	response := new(rest.CreateUserResponse)
	response.Success = true
	client.On("CreateUser", &request).Return(response, nil)

	err := c.CreateUser(&createUserVM)
	assert.NoError(t, err)

	client.AssertExpectations(t)
}

func TestCreateUserFailedLogin(t *testing.T) {
	c, client, uc := getTestableApiController()

	expectedError := errors.New("login failed test")
	client.On("Login", uc).Return(expectedError)

	err := c.CreateUser(&createUserVM)
	assert.EqualError(t, err, expectedError.Error())

	client.AssertExpectations(t)
}

func TestCreateUserFailedCreate(t *testing.T) {
	c, client, uc := getTestableApiController()

	expectedError := errors.New("create failed test")
	client.On("Login", uc).Return(nil)
	client.On("CreateUser", &request).Return(nil, expectedError)

	err := c.CreateUser(&createUserVM)
	assert.EqualError(t, err, expectedError.Error())

	client.AssertExpectations(t)
}

func TestCreateUserNoSuccess(t *testing.T) {
	c, client, uc := getTestableApiController()

	client.On("Login", uc).Return(nil)
	response := new(rest.CreateUserResponse)
	response.Success = false
	response.Error = "Test Error"
	client.On("CreateUser", &request).Return(response, nil)

	err := c.CreateUser(&createUserVM)
	assert.Error(t, err)

	client.AssertExpectations(t)
}

func TestCreateUserIgnoreExisting(t *testing.T) {
	c, client, uc := getTestableApiController()

	client.On("Login", uc).Return(nil)
	requestErr := fmt.Errorf("%s is already in use :( [error-field-unavailable]", createUserVM.Username)
	client.On("CreateUser", &request).Return(nil, requestErr)

	vm := models.CreateUserViewModel(createUserVM)
	vm.IgnoreExisting = true

	err := c.CreateUser(&vm)
	assert.NoError(t, err)

	client.AssertExpectations(t)
}

func TestPing(t *testing.T) {
	c, client, uc := getTestableApiController()
	client.On("Login", uc).Return(nil)

	err := c.Ping(1, 1*time.Microsecond, true)
	assert.NoError(t, err)

	client.AssertExpectations(t)
}

func TestPingRetry(t *testing.T) {
	c, client, uc := getTestableApiController()

	expectedError := errors.New("no access")
	client.On("Login", uc).Return(expectedError).Times(5)

	err := c.Ping(5, 1*time.Millisecond, true)
	assert.EqualError(t, err, expectedError.Error())

	client.AssertExpectations(t)
}

func TestUpdatePermissions(t *testing.T) {
	c, client, uc := getTestableApiController()

	permissions := []sdkModels.Permission{{ID: "add-user-to-any-p-room", Roles: []string{"admin", "user"}}}
	request := rest.UpdatePermissionsRequest{
		Permissions: permissions,
	}

	response := rest.UpdatePermissionsResponse{
		Status:      rest.Status{Success: true},
		Permissions: permissions,
	}

	client.On("Login", uc).Return(nil)
	client.On("UpdatePermissions", &request).Return(&response, nil)

	vm := models.UpdatePermissionsViewModel{
		PermissionId: request.Permissions[0].ID,
		Roles:        request.Permissions[0].Roles,
	}

	err := c.UpdatePermissions(&vm)
	assert.NoError(t, err)

	client.AssertExpectations(t)
}

func TestUpdatePermissionsFailedLogin(t *testing.T) {
	c, client, uc := getTestableApiController()

	expectedError := errors.New("login failed test")
	client.On("Login", uc).Return(expectedError)

	vm := models.UpdatePermissionsViewModel{}

	err := c.UpdatePermissions(&vm)
	assert.EqualError(t, err, expectedError.Error())

	client.AssertExpectations(t)
}

func TestUpdatePermissionsFailedUpdateRequest(t *testing.T) {
	c, client, uc := getTestableApiController()

	expectedError := errors.New("create failed test")

	permissions := []sdkModels.Permission{{ID: "add-user-to-any-p-room", Roles: []string{"admin", "user"}}}
	request := rest.UpdatePermissionsRequest{
		Permissions: permissions,
	}

	client.On("Login", uc).Return(nil)
	client.On("UpdatePermissions", &request).Return(nil, expectedError)

	vm := models.UpdatePermissionsViewModel{
		PermissionId: request.Permissions[0].ID,
		Roles:        request.Permissions[0].Roles,
	}

	err := c.UpdatePermissions(&vm)
	assert.EqualError(t, err, expectedError.Error())

	client.AssertExpectations(t)
}

func TestUpdatePermissionsFailedUpdateCall(t *testing.T) {
	c, client, uc := getTestableApiController()

	expectedError := errors.New("create failed test")

	permissions := []sdkModels.Permission{{ID: "add-user-to-any-p-room", Roles: []string{"admin", "user"}}}
	request := rest.UpdatePermissionsRequest{
		Permissions: permissions,
	}

	response := rest.UpdatePermissionsResponse{
		Status:      rest.Status{Success: false, Error: expectedError.Error()},
		Permissions: []sdkModels.Permission{},
	}

	client.On("Login", uc).Return(nil)
	client.On("UpdatePermissions", &request).Return(&response, nil)

	vm := models.UpdatePermissionsViewModel{
		PermissionId: request.Permissions[0].ID,
		Roles:        request.Permissions[0].Roles,
	}

	err := c.UpdatePermissions(&vm)
	assert.EqualError(t, err, expectedError.Error())

	client.AssertExpectations(t)
}
