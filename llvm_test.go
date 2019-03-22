package main

import (
	"log"
	"testing"
	"wasmgo/emscripten"
	"wasmgo/types"
)

func TestVmRun(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var vm types.VM = &emscripten.EMVM{}
	vm.Init()
	p := vm.LoadExecFile("/Users/shang/Documents/demo/a.out.wasm")
	//p := vm.LoadExecFile("/Users/shang/Desktop/code/CoinIDLib/build/core.wasm")
	vm.InvokeMethod(p, "_main")
}
