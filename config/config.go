package config

import "github.com/spf13/viper"

type Config struct {
	PORT            string `mapstructure:"PORT"`
	CQL_KEYSPACE    string `mapstructure:"CQL_KEYSPACE"`
	CQL_HOSTS       string `mapstructure:"CQL_HOSTS"`
	SPACES_ENDPOINT string `mapstructure:"SPACES_ENDPOINT"`
	SPACES_KEY      string `mapstructure:"SPACES_TOKEN"`
	SPACES_SECRET   string `mapstructure:"SPACES_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
