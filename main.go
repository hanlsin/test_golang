package main

import (
	"fmt"
	"os"

	"github.com/hanlsin/test_golang/function_ex"
	"github.com/hanlsin/test_golang/tcpserver"
)

func main() {
	fmt.Println("Hello Golang world!")

	// Call function with a string argument.
	function_ex.Hello("World!")

	// Call funcion with a string argument, and check the return value.
	fmt.Println("	Return = ", function_ex.PrintFuncInfo("Hello World!"))
	ret, _ := function_ex.GetFunctionInfo("Hello World?!?!?!?!?!?!")
	fmt.Println(ret)

	// Call function with function pointer.
	function_ex.FunctionPointer()

	// Start TCP server
	// test: echo -n "test" | netcat localhost 4567
	// stop: echo -n "STOP SERVER" | netcat localhost 4567
	fmt.Println("-------------------------")
	tcpSrvr := tcpserver.NewTCPServer()
	err := tcpSrvr.Start()
	if err != nil {
		fmt.Println("main", err)
	}

	os.Exit(0)
}
