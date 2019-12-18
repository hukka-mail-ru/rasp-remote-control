package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

func main() {
	fmt.Println("hello world")

	// Setup our service export
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, err := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)
	if err != nil {
		fmt.Println("Error!")
	}

	fmt.Println(service.Service)

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})

	for i := 0; i < 50; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("sleep: %d\n", i)
	}
	server.Shutdown()

	/*

		defer server.Shutdown()

		// Make a channel for results and start listening

		go func() {
			for i := 0; i < 50; i++ {
				time.Sleep(1 * time.Second)
				fmt.Printf("sleep: %d\n", i)
			}
		}()
	*/
}
