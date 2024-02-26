package main

/*
#cgo LDFLAGS: ./build/libgreet.a -ldl
#include "./libgreet.h"
*/
import "C"

func main() {
	C.greet(C.CString("world"))
}
