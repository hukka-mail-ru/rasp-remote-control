package main

import (
	"syscall"

	log "github.com/sirupsen/logrus"
)

func listen(ch chan []byte, device string, config *Config) {

	log.Info("Listen ", device)
	fd, err := syscall.Open(device, syscall.O_RDONLY, 0)

	if err != nil {
		log.Error("Can't open ", device, ": ", err.Error())
		ch <- []byte(config.ThreadExitMsg)
		return
	}

	defer syscall.Close(fd)

	for {

		buffer := make([]byte, config.DeviceReadBufferSize, config.DeviceReadBufferSize)
		numread, err := syscall.Read(fd, buffer)

		if err != nil {
			log.Error("Can't read ", device, ": ", err.Error())
			break
		}

		if numread > 0 {
			log.Info("Numbytes read:", numread)
			log.Info("Buffer:", buffer)

			ch <- buffer
		}
	}

	ch <- []byte(config.ThreadExitMsg)
}
