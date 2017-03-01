package function_ex

import (
	"fmt"
	"runtime"
)

type GetFuncInfoFuncType func() (string, string, int)

func getFuncInfo() (string, string, int) {
	// ignore last result value and skip this function
	pc, file, lineno, _ := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	return file, funcname, lineno
}

func Hello(str string) {
	fmt.Println("[FUNC:Hello1] Hello ", str)
}

func PrintFuncInfo(str string) int {
	// ignore last result value
	pc, file, line, _ := runtime.Caller(0)
	funcname := runtime.FuncForPC(pc).Name()
	fmt.Printf("[FUNC] %s:%s:%d\n", file, funcname, line)
	fmt.Print(str)

	return 0
}

func GetFunctionInfo(str string) (string, bool) {
	file, funcname, lineno := getFuncInfo()
	ret := fmt.Sprintf("[%s:%s:%d] %s", file, funcname, lineno, str)
	return ret, true
}

func FunctionPointer() {
	var pfnc1 GetFuncInfoFuncType
	pfnc1 = getFuncInfo
	_, funcname, lineno := pfnc1()
	fmt.Printf("111 [%s:%d]\n", funcname, lineno)

	var pfnc2 func() (string, string, int)
	pfnc2 = getFuncInfo
	_, funcname, lineno = pfnc2()
	fmt.Printf("222 [%s:%d]\n", funcname, lineno)

	pfnc3 := getFuncInfo
	_, funcname, lineno = pfnc3()
	fmt.Printf("333 [%s:%d]\n", funcname, lineno)
}
