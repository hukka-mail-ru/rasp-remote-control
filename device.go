package main

import (
	"syscall"

	log "github.com/sirupsen/logrus"
)

func listen(ch chan []byte, device string, exitMsg string) {

	log.Info("Listen ", device)

	for {
		var fd, numread int
		var err error

		fd, err = syscall.Open(device, syscall.O_RDONLY, 0777)

		if err != nil {
			log.Error("Cant open ", device, ": ", err.Error())
			ch <- []byte(exitMsg)
			return
		}

		buffer := make([]byte, 10, 100)

		numread, err = syscall.Read(fd, buffer)

		/*if err != nil {
			log.Error(err.Error(), "\n")
		}*/

		if numread > 0 {
			log.Info("Numbytes read:", numread)
			log.Info("Buffer:", buffer)

			ch <- buffer
		}

		err = syscall.Close(fd)

		if err != nil {
			log.Error(err.Error(), "\n")
		}
	}
}
