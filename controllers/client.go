package controllers

import (
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
)

type RocketChatClient interface {
	Login(credentials *models.UserCredentials) error
	CreateUser(req *models.CreateUserRequest) (*rest.CreateUserResponse, error)
	UpdatePermissions(req *rest.UpdatePermissionsRequest) (*rest.UpdatePermissionsResponse, error)
}
