package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	RfcommDevice  string
	ThreadExitMsg string
}

func getConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath("config") // directory
	var config Config

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {

	config, err := getConfig()
	if err != nil {
		log.WithError(err).Fatal("Could not load config")
		return
	}

	c1 := make(chan []byte)

	go rfcomm(c1, config)

	// MAIN LOOP
	for {
		select {
		case msg := <-c1:
			log.Info("received", msg)
			if string(msg) == config.ThreadExitMsg {
				log.Info("Exit.")
				return
			}
		}
	}

}
