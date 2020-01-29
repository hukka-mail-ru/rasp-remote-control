package rpi

import (
	"rasp-remote-control/config"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func Listen(ch chan []byte, device string) {

	log.Info("Listen ", device)
	fd, err := syscall.Open(device, syscall.O_RDONLY, 0)

	if err != nil {
		log.Error("Can't open ", device, ": ", err.Error())
		ch <- []byte(config.Inst().ThreadExitMsg)
		return
	}

	defer syscall.Close(fd)

	for {

		buffer := make([]byte, config.Inst().DeviceReadBufferSize, config.Inst().DeviceReadBufferSize)
		numread, err := syscall.Read(fd, buffer)

		if err != nil {
			log.Error("Can't read ", device, ": ", err.Error())
			break
		}

		if numread > 0 {
			//	log.Info("Numbytes read:", numread)

			readBytes := make([]byte, numread)
			copy(readBytes, buffer)

			ch <- readBytes
		}
	}

	ch <- []byte(config.Inst().ThreadExitMsg)
}
