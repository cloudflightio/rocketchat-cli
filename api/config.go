package api

import (
	"strings"
)

type Config interface {
	SetConfigType(string)
	SetEnvPrefix(string)
	SetConfigFile(string)
	AddConfigPath(string)
	SetConfigName(string)
	SetEnvKeyReplacer(*strings.Replacer)
	AutomaticEnv()
	ReadInConfig() error
	IsSet(string) bool
	GetString(string) string
	Set(key string, value interface{})
}
