package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func execTCPClient() {
	fmt.Println("[CLIENT] Execute TCP Client")

	cmd := exec.Command("netcat", "localhost", "4567")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error: stdin:", err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error: stdout:", err)
	}
	defer stdout.Close()

	reader := bufio.NewReader(stdout)
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error: cmd.Start:", err)
	}

	msgList := []string{"123", "abc", "quit", "QUIT"}
	for _, msg := range msgList {
		stdin.Write([]byte(msg))
		lineBytes, _, _ := reader.ReadLine()
		fmt.Println("[CLIENT] Response msg:", string(lineBytes))
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error: cmd.Wait:", err)
	}
}

func execTCPClient2() {
	fmt.Println("[CLIENT] Execute TCP Client2")

	cmd := exec.Command("netcat", "localhost", "4567")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error: stdin:", err)
	}
	defer stdin.Close()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error: stdout:", err)
	}
	defer stdout.Close()

	scanner := bufio.NewScanner(stdout)
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error: cmd.Start:", err)
	}

	msgList := []string{"XYZ", "9999", "stop server", "STOP SERVER"}
	for _, msg := range msgList {
		io.WriteString(stdin, msg)
		io.WriteString(stdin, "\n")
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("[CLIENT] Response msg:", line)
			break
		}
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error: cmd.Wait:", err)
	}
}

func main() {
	fmt.Println("Hello Golang world!")

	// Start client
	go func() {
		time.Sleep(time.Second * 1)
		execTCPClient()
		execTCPClient2()
	}()

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
