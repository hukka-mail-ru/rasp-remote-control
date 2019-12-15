package main

import (
	"fmt"
	"syscall"
)

func main() {

	c1 := make(chan []byte)

	go func() {
		for {
			disk := "/dev/rfcomm0"
			var fd, numread int
			var err error

			fd, err = syscall.Open(disk, syscall.O_RDONLY, 0777)

			if err != nil {
				fmt.Print(err.Error(), "\n")
				return
			}

			buffer := make([]byte, 10, 100)

			numread, err = syscall.Read(fd, buffer)

			if err != nil {
				fmt.Print(err.Error(), "\n")
			}

			fmt.Printf("Numbytes read: %d\n", numread)
			fmt.Printf("Buffer: %b\n", buffer)

			c1 <- buffer

			err = syscall.Close(fd)

			if err != nil {
				fmt.Print(err.Error(), "\n")
			}
		}
	}()

	// MAIN LOOP
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		}
	}

}
