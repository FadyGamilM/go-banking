package utils

import "github.com/spf13/viper"

type DevConfig struct {
	DbHost     string `mapstructure:"HOST"`
	DbPort     int    `mapstructure:"PORT"`
	DbUsername string `mapstructure:"USER"`
	DbPassword string `mapstructure:"PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
}

func LoadConfig(filePath string) (config DevConfig, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(filePath)

	// we need viper to also read from env vars and override those in file if we provided env vars
	viper.AutomaticEnv()

	// read the variables
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Unmarshal the config file into the DevConfig struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
