package test

import (
	"github.com/stretchr/testify/mock"
)

type MockedConfigController struct {
	mock.Mock
}

func (m *MockedConfigController) IsSet(s string) bool {
	args := m.Called(s)
	return args.Bool(0)
}

func (m *MockedConfigController) GetString(s string) string {
	args := m.Called(s)
	return args.String(0)
}
