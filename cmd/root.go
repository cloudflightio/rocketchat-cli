package cmd

import (
	"fmt"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/spf13/cobra"
	url2 "net/url"
)

var ConfigControllerFactory controllers.NewConfigController
var ApiControllerFactory controllers.NewApiController

var Verbose bool
var Config controllers.ConfigController

var CfgFile string

var apiController controllers.ApiController

var RootCmd = &cobra.Command{
	Use:   "rocketchat-cli",
	Short: "Commandline Interface for RochetChat",
	Long:  `This tool provides a basic rocketchat-api cli.`,
}

func init() {
	cobra.OnInitialize(InitConfig)

	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/.rocketchat-cli.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func InitConfig() {
	Config = ConfigControllerFactory(CfgFile, Verbose)

	if !Config.IsSet("rocketchat.url") {
		panic("config error - rocketchat.url not set")
	}

	credentials := models.UserCredentials{
		Email:    Config.GetString("user.email"),
		Token:    Config.GetString("user.token"),
		Password: Config.GetString("user.password"),
		ID:       Config.GetString("user.id"),
	}

	sUrl := Config.GetString("rocketchat.url")
	url, err := url2.Parse(sUrl)
	if err != nil {
		panic(fmt.Errorf("Fatal error: %s \n", err))
	}
	if !url.IsAbs() {
		panic(fmt.Errorf("Fatal error: %s \n", "relative url! Please provide a absolute url for the target rocketchat instance"))
	}

	apiController = ApiControllerFactory(url, Verbose, &credentials)
}
