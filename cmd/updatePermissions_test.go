package cmd

import (
	"bytes"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/mriedmann/rocketchat-cli/test"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestUpdatePermissionsCli(t *testing.T) {
	permissionsId := "add-user-to-any-p-room"
	roles := []string{"admin", "user"}

	vm := models.UpdatePermissionsViewModel{PermissionId: permissionsId, Roles: roles}
	c := test.MockedApiController{}

	ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		c.On("UpdatePermissions", &vm).Return(nil)
		return &c
	}
	ConfigControllerFactory = NewMockedConfigController

	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs([]string{"updatePermissions", "-d", permissionsId, "-r", strings.Join(roles, ",")})

	cmd, err := rootCmd.ExecuteC()
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cmd)
	c.AssertExpectations(t)
}
