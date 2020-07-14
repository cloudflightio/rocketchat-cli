package cmd

import (
	"bytes"
	"github.com/cloudflightio/rocketchat-cli/cmd"
	"github.com/cloudflightio/rocketchat-cli/controllers"
	"github.com/cloudflightio/rocketchat-cli/models"
	"github.com/cloudflightio/rocketchat-cli/test"
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

	cmd.ApiControllerFactory = func(url *url.URL, b bool, credentials *models.UserCredentials) controllers.ApiController {
		c.On("UpdatePermissions", &vm).Return(nil)
		return &c
	}
	cmd.ConfigControllerFactory = NewMockedConfigController

	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOutput(buf)
	cmd.RootCmd.SetArgs([]string{"updatePermissions", "-d", permissionsId, "-r", strings.Join(roles, ",")})

	cli, err := cmd.RootCmd.ExecuteC()
	assert.NoError(t, err)

	output := buf.String()
	assert.Empty(t, output)

	assert.NotNil(t, cli)
	c.AssertExpectations(t)
}
