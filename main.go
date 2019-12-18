package main

import (
	"bytes"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	RfcommDevice string
}

func getConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("common")
	v.AddConfigPath("config")
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
			if bytes.Compare(msg, ExitMsg) == 0 {
				log.Info("Exit.")
				return
			}
		}
	}

}
