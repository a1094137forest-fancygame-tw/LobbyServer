package constant

import (
	"log"

	"github.com/spf13/viper"

	"LobbyServer/config"
)

func ReadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")

	viper.SetDefault("RUN_MODE", "debug")

	envs := []string{
		"LOCAL_SERVER_URL",
		"LOBBY_SERVER_URL",
		"USER_SERVER_URL",
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			log.Println(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	config.Initialize()
}
