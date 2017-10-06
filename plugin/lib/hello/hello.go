package hello

import (
	"fmt"
)

// Arg string variable
var Arg string

// Args 10 strings array variable
var Args [10]string

// ArgNum int variable
var ArgNum int

// PrintHello print hello
func PrintHello() {
	fmt.Println("Hello")
}

// PrintArg1 print hello with argument
func PrintArg1(arg string) {
	Arg = arg
	ArgNum = 1
	fmt.Println("Hello ", Arg)
}
