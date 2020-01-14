package main

import (
	"syscall"

	log "github.com/sirupsen/logrus"
)

func listen(ch chan []byte, device string, exitMsg string) {

	log.Info("Listen ", device)

	var fd, numread int
	var err error

	//fd, err = syscall.Open(device, syscall.O_RDONLY, 0)
	fd, err = syscall.Open(device, syscall.O_RDONLY, 0)

	if err != nil {
		log.Error("Can't open ", device, ": ", err.Error())
		ch <- []byte(exitMsg)
		return
	}

	defer syscall.Close(fd)

	for {

		buffer := make([]byte, 1000, 1000)

		numread, err = syscall.Read(fd, buffer)

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

	ch <- []byte(exitMsg)
}
