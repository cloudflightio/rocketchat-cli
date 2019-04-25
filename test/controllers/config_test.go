package controllers

import (
	"github.com/magiconair/properties/assert"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"os"
	"testing"
)

func TestNewViperConfigController(t *testing.T) {
	url := "https://test:1234"
	_ = os.Setenv("RCCLI_ROCKETCHAT_URL", url)

	c := controllers.NewViperConfigController("", true)

	assert.Equal(t, c.GetString("rocketchat.url"), url)
	_ = os.Unsetenv("RCCLI_ROCKETCHAT_URL")
}

func TestNewViperConfigControllerFile(t *testing.T) {
	file := "../../config.example.yaml"

	c := controllers.NewViperConfigController(file, true)

	assert.Equal(t, c.GetString("rocketchat.url"), "http://localhost:3000")
}
