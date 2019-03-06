package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"wasmgo/emscripten"
	"wasmgo/types"
)

//go:generate go build -buildmode=plugin /Users/shang/go/src/wasmgo/plugin/network
//go:generate go build
//go:generate go build -buildmode=c-archive -o /Users/shang/Documents/wasm-java/lib/wasm.a

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
		result := vm.InvokeMethod(p, *name, args...)
		fmt.Printf("result= %d \n", result)
	} else {
		flag.PrintDefaults()
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("参数错误")
		}
	}()
}
