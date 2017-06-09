package main

import (
	"fmt"
	"os"
)

func execTCPClient() {
}

func main() {
	fmt.Println("Hello Golang world!")

	// Start TCP server
	// test: echo -n "test" | netcat localhost 4567
	// stop: echo -n "STOP SERVER" | netcat localhost 4567
	fmt.Println("-------------------------")
	tcpSrvr := NewTCPServer()
	err := tcpSrvr.Start()
	if err != nil {
		fmt.Println("main", err)
	}

	os.Exit(0)
}
