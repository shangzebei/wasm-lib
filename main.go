package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"wasmgo/llvm"
)

func main() {
	log.SetOutput(ioutil.Discard)
	arg := os.Args
	_, err := os.Stat(arg[1])
	if err == nil {
		llvm.CallMain(arg[1])
	} else {
		fmt.Printf("file %s err ", arg[1])
	}

}
