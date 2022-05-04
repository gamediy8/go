package utils

import (
	"github.com/spf13/viper"
	"log"
)

var Config *viper.Viper

func GetConfig() *viper.Viper {
	if Config == nil {
		viper.SetConfigName("config")    // name of config file (without extension)
		viper.SetConfigType("toml")      // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("./config/") // path to look for the config file in
		// optionally look for config in the working directory
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {

			log.Println()
			panic(err)
		}
		Config = viper.GetViper()
	}
	return Config

}
