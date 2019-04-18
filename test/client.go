package test

import (
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
	"github.com/stretchr/testify/mock"
)

type MockedClient struct {
	mock.Mock
}

func (m *MockedClient) Login(credentials *models.UserCredentials) error {
	args := m.Called(credentials)
	return args.Error(0)
}

func (m *MockedClient) CreateUser(req *models.CreateUserRequest) (*rest.CreateUserResponse, error) {
	args := m.Called(req)
	if r := args.Get(0); r == nil {
		return nil, args.Error(1)
	} else {
		return r.(*rest.CreateUserResponse), args.Error(1)
	}
}

func (m *MockedClient) UpdatePermissions(req *rest.UpdatePermissionsRequest) (*rest.UpdatePermissionsResponse, error) {
	args := m.Called(req)
	if r := args.Get(0); r == nil {
		return nil, args.Error(1)
	} else {
		return r.(*rest.UpdatePermissionsResponse), args.Error(1)
	}
}
