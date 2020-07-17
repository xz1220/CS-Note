package main

//#include "hello.h"
import "C"
import "fmt"

func main() {
	fmt.Println("Test")
	C.sayHello(C.CString("Hello\n"))
}
