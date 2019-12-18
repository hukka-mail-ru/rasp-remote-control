package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio"
)

func enablePin(pinNumber int) {
	log.Info("Enable pin ", pinNumber)
	rpio.Open()
	defer rpio.Close()

	pin := rpio.Pin(pinNumber)
	pin.Output()
	pin.High()
}

func disablePin(pinNumber int) {
	rpio.Open()
	defer rpio.Close()

	pin := rpio.Pin(pinNumber)
	pin.Output()
	pin.Low()
}
