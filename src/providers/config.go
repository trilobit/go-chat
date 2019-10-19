package providers

import (
	"github.com/spf13/viper"
)

// NewConfig creates instance of config
func NewConfig() (*viper.Viper, error) {
	cnf := viper.New()
	cnf.AddConfigPath(".")
	cnf.SetConfigFile("config.yml")
	cnf.SetConfigType("yaml")
	cnf.AutomaticEnv()

	if err := cnf.ReadInConfig(); err != nil {
		return nil, err
	}

	return cnf, nil
}
