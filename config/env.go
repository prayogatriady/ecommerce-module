package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewConfig(appEnv string, dir string) error {

	switch appEnv {
	case "PROD":
		viper.SetConfigName("config-prod")
	case "STG":
		viper.SetConfigName("config-stg")
	case "DEV":
		viper.SetConfigName("config-dev")
	default:
		viper.SetConfigName("config")
	}

	viper.SetConfigType("json")
	viper.AddConfigPath(dir)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("couldn't read config: %v", err)
	}

	return nil
}

func String(key string, def string) string {
	val := viper.GetString(key)
	if val == "" {
		return def
	}
	return val
}

func Int(key string, def int) int {
	val := viper.GetInt(key)
	if val == 0 {
		return def
	}
	return val
}

func Bool(key string) bool {
	return viper.GetBool(key)
}

func Any(key string, def any) any {
	val := viper.Get(key)
	if val == nil {
		return def
	}
	return val
}
