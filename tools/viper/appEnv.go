package viper

import (
	"github.com/spf13/viper"
)

// ViperGetEnv using viper for getting env
func ViperGetEnv(key, defaultValue string) string {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
