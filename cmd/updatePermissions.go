package cmd

import (
	"github.com/cloudflightio/rocketchat-cli/models"
	"github.com/spf13/cobra"
)

var updatePermissions = &cobra.Command{
	Use:   "updatePermissions",
	Short: "updated permissions using the rest-api. Proper rights needed.",
	Long:  ``,
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
	RootCmd.AddCommand(updatePermissions)

	flags := updatePermissions.PersistentFlags()
	flags.StringVarP(&PermissionId, "id", "d", "", "Target permission id (e.g. add-user-to-any-p-room)")
	flags.StringSliceVarP(&Roles, "roles", "r", []string{}, "Roles that should have the given permission")

	_ = cobra.MarkFlagRequired(flags, "id")
	_ = cobra.MarkFlagRequired(flags, "roles")
}
