package controllers

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type NewConfigController func(string, bool) ConfigController

type ConfigController interface {
	IsSet(string) bool
	GetString(string) string
}

func NewViperConfigController(cfgFile string, verbose bool) ConfigController {
	c := viper.New()
	c.SetConfigType("yaml")
	c.SetEnvPrefix("rccli")

	if cfgFile != "" {
		c.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, _ := homedir.Dir()

		// Search config in home directory with name ".rocketchat-cli" (without extension).
		c.AddConfigPath(home)
		c.AddConfigPath(".")
		c.SetConfigName(".rocketchat-cli")
	}

	c.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := c.ReadInConfig(); err == nil {
		if verbose {
			log.Printf("Using config file: %s \n", c.ConfigFileUsed())
		}
	}
	return c
}
