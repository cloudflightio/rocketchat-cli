// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "createUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vm := models.CreateUserViewModel{
			Name:           NewUserName,
			Email:          NewUserEmail,
			Password:       NewUserPassword,
			Username:       NewUserUsername,
			IgnoreExisting: IgnoreExisting,
		}
		return apiController.CreateUser(&vm)
	},
}

var NewUserName string
var NewUserEmail string
var NewUserPassword string
var NewUserUsername string

var IgnoreExisting bool

func init() {
	rootCmd.AddCommand(createUserCmd)

	flags := createUserCmd.PersistentFlags()
	flags.StringVarP(&NewUserUsername, "username", "u", "", "Target username of newly created user")
	flags.StringVarP(&NewUserEmail, "email", "e", "", "Target email address of newly created user")

	flags.StringVarP(&NewUserPassword, "password", "p", "", "Target password of newly created user")

	flags.StringVarP(&NewUserName, "name", "n", "", "Target real name of newly created user")
	flags.BoolVarP(&IgnoreExisting, "ignore-existing", "i", false, "Continue without error if the given target user already exists")

	_ = cobra.MarkFlagRequired(flags, "username")
	_ = cobra.MarkFlagRequired(flags, "username")
	_ = cobra.MarkFlagRequired(flags, "email")
	_ = cobra.MarkFlagRequired(flags, "password")
}
