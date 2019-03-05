package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"wasmgo/emscripten"
	"wasmgo/types"
)

func main() {
	name := flag.String("n", "_main", "wasm method name")
	file := flag.String("f", "", "wasm file name ")
	debug := flag.Bool("v", false, "open debug mode (-v open)")
	flag.Parse()

	args := flag.Args()

	if *debug != true {
		log.SetOutput(ioutil.Discard)
	}

	if *file != "" {
		var vm types.VM = &emscripten.EMVM{}
		vm.Init()
		p := vm.LoadExecFile(*file)
		vm.InvokeMethod(p, *name, args...)
	} else {
		flag.PrintDefaults()
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("参数错误")
		}
	}()
}
