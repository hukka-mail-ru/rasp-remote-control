package main

import (
	"fmt"
	"syscall"
)

func main() {
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

	err = syscall.Close(fd)

	if err != nil {
		fmt.Print(err.Error(), "\n")
	}
}
