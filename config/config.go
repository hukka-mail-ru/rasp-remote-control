package config

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	LircDevice           string
	RfcommDevice         string
	DeviceReadBufferSize int
	ThreadExitMsg        string
	PinNumber            int
}

var instance *Config
var once sync.Once

// singleton
func Inst() *Config {
	once.Do(func() {
		instance = &Config{}

		// config
		v := viper.New()

		v.SetConfigName("config")
		v.AddConfigPath("config") // directory

		err := v.ReadInConfig()
		if err != nil {
			log.Error("Can't load config", err)
			return
		}

		err = v.Unmarshal(instance)
		if err != nil {
			log.Error("Can't load config", err)
			return
		}

	})
	return instance
}
