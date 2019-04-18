package cmd

import (
	"bytes"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/mriedmann/rocketchat-cli/test"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestUpdatePermissionsCli(t *testing.T) {
	Config = viper.New()
	Config.Set("rocketchat.url", "http://localhost:3000")

	permissionsId := "add-user-to-any-p-room"
	roles := []string{"admin", "user"}

	vm := models.UpdatePermissionsViewModel{PermissionId: permissionsId, Roles: roles}

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		apiController := test.MockedApiController{}
		apiController.On("UpdatePermissions", &vm).Return(nil)
		return &apiController
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs([]string{"updatePermissions", "-d", permissionsId, "-r", strings.Join(roles, ",")})

	cmd, err := rootCmd.ExecuteC()
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cmd)
	apiController.(*test.MockedApiController).AssertExpectations(t)
}
