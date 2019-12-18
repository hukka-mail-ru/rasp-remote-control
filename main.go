package main

import (
	"bytes"

	log "github.com/sirupsen/logrus"
)

func main() {

	c1 := make(chan []byte)

	go rfcomm(c1)

	// MAIN LOOP
	for {
		select {
		case msg := <-c1:
			log.Info("received", msg)
			if bytes.Compare(msg, []byte{01, 02, 03}) == 0 {
				log.Info("Exit.")
				return
			}
		}
	}

}
