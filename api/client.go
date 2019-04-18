package api

import (
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
	"net/url"
)

type RocketChatClient interface {
	Login(credentials *models.UserCredentials) error
	CreateUser(req *models.CreateUserRequest) (*rest.CreateUserResponse, error)
	UpdatePermissions(req *rest.UpdatePermissionsRequest) (*rest.UpdatePermissionsResponse, error)
	Get(api string, params url.Values, response rest.Response) error
}
