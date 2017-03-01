package main

import "fmt"
import "os"

func main() {
	fmt.Println("Hello Golang world!")

	// Call function with a string argument.
	hello("World!")

	// Call funcion with a string argument, and check the return value.
	fmt.Println("	Return = ", printFuncInfo("Hello World!"))
	ret, _ := getFunctionInfo("Hello World?!?!?!?!?!?!")
	fmt.Println(ret)

	// Call function with function pointer.
	exFunctionPointer()

	os.Exit(0)
}
