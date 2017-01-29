package config

import "github.com/spf13/viper"

func Config() {
	viper.AutomaticEnv()

	viper.SetDefault("env", "dev")

	viper.SetConfigName("config")

	viper.AddConfigPath("config/")
	viper.ReadInConfig()
}

func HasToken(token string) bool {
	tokens := viper.GetStringSlice("tokens")
	for _, a := range tokens {
		if a == token {
			return true
		}
	}
	return false
}
