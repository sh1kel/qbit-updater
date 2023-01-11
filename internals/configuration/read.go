package configuration

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func (c Config) String() string {
	return fmt.Sprintf("Clients addresses: %s",
		c.Clients.Urls)
}

func InitConfig(configFile string) *Config {
	newConfig := new(Config)
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".qbit-updater")
	}
	if err := viper.ReadInConfig(); err == nil {
		log.Info("using configuration file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&newConfig); err != nil {
		log.Fatal(err)
	}
	log.Debug(newConfig)
	return newConfig
}
