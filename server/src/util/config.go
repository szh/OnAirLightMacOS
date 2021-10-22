package util

import "github.com/spf13/viper"

type ConfigStruct struct {
	Port     string `mapstructure:"PORT"`
	PinRed   int    `mapstructure:"PIN_RED"`
	PinGreen int    `mapstructure:"PIN_GREEN"`
	TestMode bool   `mapstructure:"TEST_MODE"`
}

var Config ConfigStruct

func LoadConfig(path string) (config ConfigStruct, err error) {
	// config file
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// environment variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	Config = config
	return
}
