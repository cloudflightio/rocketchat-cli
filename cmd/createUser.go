package cmd

import (
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "createUser",
	Short: "Creates a new user, using the rest api",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		vm := models.CreateUserViewModel{
			Name:           NewUserName,
			Email:          NewUserEmail,
			Password:       NewUserPassword,
			Username:       NewUserUsername,
			Roles:          NewUserRoles,
			IgnoreExisting: IgnoreExisting,
		}
		return apiController.CreateUser(&vm)
	},
}

var NewUserName string
var NewUserEmail string
var NewUserPassword string
var NewUserUsername string
var NewUserRoles []string

var IgnoreExisting bool

func init() {
	RootCmd.AddCommand(createUserCmd)

	flags := createUserCmd.PersistentFlags()
	flags.StringVarP(&NewUserUsername, "username", "u", "", "Target username of newly created user")
	flags.StringVarP(&NewUserEmail, "email", "e", "", "Target email address of newly created user")
	flags.StringVarP(&NewUserPassword, "password", "p", "", "Target password of newly created user")
	flags.StringVarP(&NewUserName, "name", "n", "", "Target real name of newly created user")
	flags.StringSliceVarP(&NewUserRoles, "roles", "r", []string(nil), "Target roles of newly created user")
	flags.BoolVarP(&IgnoreExisting, "ignore-existing", "i", false, "Continue without error if the given target user already exists")

	_ = cobra.MarkFlagRequired(flags, "username")
	_ = cobra.MarkFlagRequired(flags, "username")
	_ = cobra.MarkFlagRequired(flags, "email")
	_ = cobra.MarkFlagRequired(flags, "password")
}
