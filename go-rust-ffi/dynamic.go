package main

/*
#cgo LDFLAGS: -L./build -lgreet
#include "./libgreet.h"
*/
import "C"

func main() {
	C.greet(C.CString("world"))
}
