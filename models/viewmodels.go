package models

type UserCredentials struct {
	Email    string
	Password string
	Token    string
	ID       string
}

type CreateUserViewModel struct {
	Name           string
	Email          string
	Password       string
	Username       string
	IgnoreExisting bool
}

type UpdatePermissionsViewModel struct {
	PermissionId string
	Roles        []string
}
