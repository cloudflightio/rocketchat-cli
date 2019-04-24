package controllers

import (
	"errors"
	sdkModels "github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
	"github.com/matryer/try"
	"github.com/mriedmann/rocketchat-cli/api"
	"github.com/mriedmann/rocketchat-cli/models"
	"log"
	"net/url"
	"strings"
	"time"
)

type NewApiController func(*url.URL, bool, *models.UserCredentials) ApiController

type ApiController interface {
	CreateUser(*models.CreateUserViewModel) error
	Ping(int, time.Duration, bool) error
	UpdatePermissions(*models.UpdatePermissionsViewModel) error
}

type SdkApiController struct {
	client      api.RocketChatClient
	credentials *sdkModels.UserCredentials
}

func NewSdkApiController(serverUrl *url.URL, debug bool, credentials *models.UserCredentials) (c ApiController) {
	uc := sdkModels.UserCredentials{
		Email:    credentials.Email,
		Password: credentials.Password,
		ID:       credentials.ID,
		Token:    credentials.Token,
	}
	return &SdkApiController{
		client:      rest.NewClient(serverUrl, debug),
		credentials: &uc,
	}
}

func (c *SdkApiController) login() (err error) {
	err = c.client.Login(c.credentials)
	return
}

func (c *SdkApiController) CreateUser(model *models.CreateUserViewModel) (err error) {
	err = c.login()
	if err != nil {
		return
	}

	request := sdkModels.CreateUserRequest{
		Name:         model.Name,
		Email:        model.Email,
		Password:     model.Password,
		Username:     model.Username,
		CustomFields: nil,
	}

	response, err := c.client.CreateUser(&request)
	if err != nil || !response.Success {
		if err == nil {
			err = errors.New(response.Error)
		}
		if model.IgnoreExisting && strings.HasSuffix(err.Error(), "[error-field-unavailable]") {
			err = nil
		} else {
			return
		}
	}

	return
}

func (c *SdkApiController) Ping(maxAttempts int, waitTime time.Duration, verbose bool) error {
	err := try.Do(func(attempt int) (bool, error) {
		err := c.login()
		if err != nil {
			if verbose {
				log.Printf("error (attempt %d): %s \n", attempt, err)
			}
			time.Sleep(waitTime)
		}
		return attempt < maxAttempts, err
	})

	return err
}

func (c *SdkApiController) UpdatePermissions(model *models.UpdatePermissionsViewModel) (err error) {
	err = c.login()
	if err != nil {
		return
	}

	request := rest.UpdatePermissionsRequest{
		Permissions: []sdkModels.Permission{{ID: model.PermissionId, Roles: model.Roles}},
	}
	response, err := c.client.UpdatePermissions(&request)
	if err != nil {
		return
	}

	if !response.Success {
		err = errors.New(response.Error)
		return
	}

	return
}
