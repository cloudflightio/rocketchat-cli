package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mriedmann/rocketchat-cli/api"
	"github.com/mriedmann/rocketchat-cli/controllers"
	"github.com/mriedmann/rocketchat-cli/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	url2 "net/url"
	"strings"
)

var ApiControllerFactory controllers.NewApiController

var Verbose bool
var Config api.Config

var cfgFile string

var apiController controllers.ApiController

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rocketchat-cli",
	Short: "Commandline Interface for RochetChat",
	Long: `This tool provides a basic rocketchat-api cli.
	
The main perpose `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rocketchat-cli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	Config.SetConfigType("yaml")
	Config.SetEnvPrefix("rccli")

	if cfgFile != "" {
		// Use config file from the flag.
		Config.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		// Search config in home directory with name ".rocketchat-cli" (without extension).
		Config.AddConfigPath(home)
		Config.AddConfigPath(".")
		Config.SetConfigName(".rocketchat-cli")
	}

	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	Config.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := Config.ReadInConfig(); err == nil {
		if Verbose {
			log.Printf("Using config file: %s \n", viper.ConfigFileUsed())
		}
	}

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

	apiController = ApiControllerFactory(url, Verbose, &credentials)
}
