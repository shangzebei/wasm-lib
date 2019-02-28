package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"wasmgo/emscripten"
	"wasmgo/types"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("参数错误")
		}
	}()
	log.SetOutput(ioutil.Discard)
	arg := os.Args
	_, err := os.Stat(arg[1])
	if err == nil {
		var vm types.VM = &emscripten.EMVM{}
		vm.Init()
		p := vm.LoadExecFile(arg[1])
		vm.InvokeMethod(p, "main")
	} else {
		fmt.Printf("file %s err ", arg[1])
	}

}
