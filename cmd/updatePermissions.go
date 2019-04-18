package cmd

import (
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
var updatePermissions = &cobra.Command{
	Use:   "updatePermissions",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		vm := models.UpdatePermissionsViewModel{
			PermissionId: PermissionId,
			Roles:        Roles,
		}
		err = apiController.UpdatePermissions(&vm)
		return
	},
}

var PermissionId string

var Roles []string

func init() {
	rootCmd.AddCommand(updatePermissions)

	flags := updatePermissions.PersistentFlags()
	flags.StringVarP(&PermissionId, "id", "d", "", "Target permission id (e.g. add-user-to-any-p-room)")
	flags.StringSliceVarP(&Roles, "roles", "r", []string{}, "Roles that should have the given permission")

	_ = cobra.MarkFlagRequired(flags, "id")
	_ = cobra.MarkFlagRequired(flags, "roles")
}
