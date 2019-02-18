package main

import (
	"io/ioutil"
	"log"
	"os"
	"wasmgo/llvm"
)

func main() {
	log.SetOutput(ioutil.Discard)
	arg := os.Args
	llvm.CallMain(arg[1])
}
