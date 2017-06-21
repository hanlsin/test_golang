package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Golang world!")

	// Call function with a string argument.
	Hello("World!")

	// Call funcion with a string argument, and check the return value.
	fmt.Println("	Return = ", PrintFuncInfo("Hello World!"))
	ret, _ := GetFunctionInfo("Hello World?!?!?!?!?!?!")
	fmt.Println(ret)

	// Call function with function pointer.
	FunctionPointer()

	// test ListPrintf
	LikePrintf("%s.....%d\n", "Hello?", 5)

	os.Exit(0)
}
