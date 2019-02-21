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
		p := llvm.LoadExecFile(arg[1])
		llvm.InvokeMethod(p, "main")
	} else {
		fmt.Printf("file %s err ", arg[1])
	}

}
