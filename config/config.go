package config

import "github.com/spf13/viper"

var (
	Port string

	LocalServerUrl string
	UserServerUrl  string
	LobbyServerUrl string
)

func Initialize() {
	Port = viper.GetString("PORT")

	LocalServerUrl = viper.GetString("LOCAL_SERVER_URL")
	UserServerUrl = viper.GetString("USER_SERVER_URL")
	LobbyServerUrl = viper.GetString("LOBBY_SERVER_URL")

}
