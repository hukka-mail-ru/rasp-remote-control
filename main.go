package main

import (
	"fmt"
	"os"

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
	defer server.Shutdown()

	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	// Start the lookup
	mdns.Lookup("_foobar._tcp", entriesCh)
	close(entriesCh)
}
