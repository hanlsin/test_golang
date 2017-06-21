package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// https://cryptic.io/go-http/

func execHTTPClient() {
	fmt.Println("[CLIENT] Execute HTTP Client")

	out, err := exec.Command("curl", "http://www.google.com").Output()
	if err != nil {
		fmt.Println("Error: out:", err)
	}
	fmt.Println("[CLIENT] Response msg length from http://www.google.com:", len(out))

	out, err = exec.Command("curl", "http://localhost:4567/foo").Output()
	if err != nil {
		fmt.Println("Error: out:", err)
	}
	fmt.Println("[CLIENT] Response msg:", string(out))

	out, err = exec.Command("curl", "http://localhost:4567/bar").Output()
	if err != nil {
		fmt.Println("Error: out:", err)
	}
	fmt.Println("[CLIENT] Response msg:", string(out))

	out, err = exec.Command("curl", "http://localhost:4567").Output()
	if err != nil {
		fmt.Println("Error: out:", err)
	}
	fmt.Println("[CLIENT] Response msg:", string(out))
}

func main() {
	fmt.Println("Hello Golang world!")
	fmt.Println("-------------------------")

	done := make(chan int)

	// Make HTTP Server
	httpSrvr := newHttpServer()

	// Start client
	go func(*http.Server) {
		time.Sleep(time.Second * 1)
		execHTTPClient()

		// Stop HTTP server
		err := httpSrvr.Shutdown(nil)
		if err != nil {
			fmt.Println("ERROR: Shutdown:", err)
		}
		done <- 0
	}(httpSrvr)

	// Start HTTP server
	StartHTTPServer(httpSrvr)

	<-done

	os.Exit(0)
}
