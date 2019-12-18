package main

import (
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	RfcommDevice  string
	ThreadExitMsg string
	PinNumber     int
}

func getConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath(".") // directory
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

func parseMsg(msg string) (int, error) {
	//	log.Info("parsing: ", string(msg))

	msg = strings.TrimFunc(msg, func(r rune) bool {
		return r == '\n' || r == '\x00'
	})

	//	log.Info("trimmed string: ", string(msg))

	var i int
	var err error
	i, err = strconv.Atoi(string(msg))

	return i, err
}

func main() {

	// logger
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05Z07:00",
	})
	log.Info("Hello!")

	// config
	config, err := getConfig()
	if err != nil {
		log.WithError(err).Fatal("Could not load config")
		return
	}

	// Start Rfcomm
	c1 := make(chan []byte)
	go rfcomm(c1, config)

	// MAIN LOOP
	for {
		select {
		case msg := <-c1:
			//log.Info("received string: ", string(msg))
			if string(msg) == config.ThreadExitMsg {
				log.Info("Exit.")
				return
			}

			var i int
			var err error
			i, err = parseMsg(string(msg))

			if err != nil {
				log.Error(err)
				continue
			}

			go func() {
				enablePin(config.PinNumber)
				log.Info("sleep, msec: ", i)
				time.Sleep(time.Duration(i) * time.Millisecond)
				disablePin(config.PinNumber)
			}()

		}
	}

}
