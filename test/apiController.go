package test

import (
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockedApiController struct {
	mock.Mock
}

func (m *MockedApiController) CreateUser(model *models.CreateUserViewModel) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *MockedApiController) Ping(maxAttempts int, waitTime time.Duration, verbose bool) error {
	args := m.Called(maxAttempts, waitTime, verbose)
	return args.Error(0)
}

func (m *MockedApiController) UpdatePermissions(model *models.UpdatePermissionsViewModel) error {
	args := m.Called(model)
	return args.Error(0)
}
