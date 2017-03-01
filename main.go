package main

import (
	"fmt"
	"github.com/hanlsin/test_golang/function_ex"
	"os"
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

	os.Exit(0)
}
