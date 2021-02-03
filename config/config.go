package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

func NewConfig() *Config {
	c := new(Config)
	c.config = readConfig()
	return c
}

func (c *Config) Get() *viper.Viper {
	if c.config == nil {
		log.Fatal("config not initialized")
	}
	return c.config
}

func readConfig() *viper.Viper {
	log.Info("reading environment variables")
	v := viper.New()
	v.AutomaticEnv()
	env := v.GetString("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	log.Infof("ENVIRONMENT: %s", env)
	v.SetConfigName(env)
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file or env variable '%s'", err.Error())
	}

	return v
}
