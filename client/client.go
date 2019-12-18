package main

import (
	"fmt"
	//"os"

	"github.com/hashicorp/mdns"
)

func main() {
	fmt.Println("client")

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
