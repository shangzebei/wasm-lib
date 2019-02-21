package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	log.SetOutput(ioutil.Discard)
	arg := os.Args
	_, err := os.Stat(arg[1])
	if err == nil {
		p := LoadExecFile(arg[1])
		InvokeMethod(p, "main")
	} else {
		fmt.Printf("file %s err ", arg[1])
	}

}
