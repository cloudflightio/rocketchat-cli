package controllers

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestNewViperConfigController(t *testing.T) {
	url := "https://test:1234"
	_ = os.Setenv("RCCLI_ROCKETCHAT_URL", url)

	c := NewViperConfigController("", true)

	assert.Equal(t, c.GetString("rocketchat.url"), url)
	_ = os.Unsetenv("RCCLI_ROCKETCHAT_URL")
}

func TestNewViperConfigControllerFile(t *testing.T) {
	file := "../config.example.yaml"

	c := NewViperConfigController(file, true)

	assert.Equal(t, c.GetString("rocketchat.url"), "http://localhost:3000")
}
