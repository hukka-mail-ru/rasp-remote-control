package main

import (
	"rasp-remote-control/bluetooth"
	"rasp-remote-control/config"
	"rasp-remote-control/lirc"
	"rasp-remote-control/rpi"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {

	// logger
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05Z07:00",
	})
	log.Info("Hello!")

	// Start Rfcomm
	chRfcomm := make(chan []byte)
	go rpi.Listen(chRfcomm, config.Inst().RfcommDevice)

	// Start Lirc
	chLirc := make(chan []byte)
	go rpi.Listen(chLirc, config.Inst().LircDevice)

	// MAIN LOOP
	for {
		select {
		case msg := <-chRfcomm:
			log.Info("From chRfcomm: ", string(msg))
			if string(msg) == config.Inst().ThreadExitMsg {
				log.Info("Exit.")
				return
			}

			var i int
			var err error
			i, err = bluetooth.ParseMsg(string(msg))

			if err != nil {
				log.Error(err)
				continue
			}

			go func() {
				rpi.EnablePin(config.Inst().PinNumber)
				log.Info("sleep, msec: ", i)
				time.Sleep(time.Duration(i) * time.Millisecond)
				rpi.DisablePin(config.Inst().PinNumber)
			}()

		case msg := <-chLirc:
			log.Info("From LIRC: ", msg, string(msg))

			if string(msg) == config.Inst().ThreadExitMsg {
				log.Info("Exit.")
				return
			}

			arr := lirc.ConvertToArrayOfUint32(msg)
			lirc.PrintPulseSpace(arr)

			continue
		}
	}

}
