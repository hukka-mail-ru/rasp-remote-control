package main

import (
	"strconv"
	"strings"

	"encoding/binary"
	"fmt"
	"time"

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

func convertToArrayOfUint32(bytes []byte) []uint32 {

	var array []uint32

	for i := 0; i < len(bytes)-4; i += 4 {

		res := binary.LittleEndian.Uint32(bytes[i : i+4])

		array = append(array, res)
	}

	return array
}

func printPulseSpace(values []uint32) {

	const PULSE_MASK = 0x01000000
	const TIMEOUT_MASK = 0x03000000
	const VALUE_MASK = 0x00FFFFFF

	for i := 0; i < len(values); i++ {

		fmt.Printf("Uint32: %08x ", values[i])

		pulse := (values[i] & PULSE_MASK) > 0
		timeout := (values[i] & TIMEOUT_MASK) > 0
		val := values[i] & VALUE_MASK

		if pulse {

			fmt.Printf("pulse %d", val)
			if val > 8000 {

			}

		} else if timeout {

			fmt.Printf("timeout\n")

		} else {

			fmt.Printf("space %d\n", val)
		}
	}

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
	chRfcomm := make(chan []byte)
	go listen(chRfcomm, config.RfcommDevice, config)

	// Start Lirc
	chLirc := make(chan []byte)
	go listen(chLirc, config.LircDevice, config)

	// MAIN LOOP
	for {
		select {
		case msg := <-chRfcomm:
			log.Info("From chRfcomm: ", string(msg))
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

		case msg := <-chLirc:
			log.Info("From LIRC: ", msg, string(msg))

			if string(msg) == config.ThreadExitMsg {
				log.Info("Exit.")
				return
			}

			arr := convertToArrayOfUint32(msg)
			printPulseSpace(arr)

			continue
		}
	}

}
