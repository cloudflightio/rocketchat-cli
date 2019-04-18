package test

import (
	"github.com/stretchr/testify/mock"
	"strings"
)

type MockedConfig struct {
	mock.Mock
}

func (m *MockedConfig) SetConfigType(s string) {
	m.Called(s)
}

func (m *MockedConfig) SetEnvPrefix(s string) {
	m.Called(s)
}

func (m *MockedConfig) SetConfigFile(s string) {
	m.Called(s)
}

func (m *MockedConfig) AddConfigPath(s string) {
	m.Called(s)
}

func (m *MockedConfig) SetConfigName(s string) {
	m.Called(s)
}

func (m *MockedConfig) SetEnvKeyReplacer(r *strings.Replacer) {
	m.Called(r)
}

func (m *MockedConfig) AutomaticEnv() {
	m.Called()
}

func (m *MockedConfig) ReadInConfig() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedConfig) IsSet(s string) bool {
	args := m.Called(s)
	return args.Bool(0)
}

func (m *MockedConfig) GetString(s string) string {
	args := m.Called(s)
	return args.String(0)
}

func (m *MockedConfig) Set(key string, value interface{}) {}
