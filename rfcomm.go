package main

import (
	"syscall"

	log "github.com/sirupsen/logrus"
)

func rfcomm(c1 chan []byte) {

	for {
		device := "/dev/rfcomm0"
		var fd, numread int
		var err error

		fd, err = syscall.Open(device, syscall.O_RDONLY, 0777)

		if err != nil {
			log.Error("Cant open ", device, ": ", err.Error(), "\n")
			c1 <- []byte{01, 02, 03}
			return
		}

		buffer := make([]byte, 10, 100)

		numread, err = syscall.Read(fd, buffer)

		if err != nil {
			log.Error(err.Error(), "\n")
		}

		log.Info("Numbytes read: %d\n", numread)
		log.Info("Buffer: %b\n", buffer)

		c1 <- buffer

		err = syscall.Close(fd)

		if err != nil {
			log.Error(err.Error(), "\n")
		}
	}
}
