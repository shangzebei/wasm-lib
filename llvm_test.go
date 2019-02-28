package main

import (
	"testing"
	"wasmgo/emscripten"
	"wasmgo/types"
)

func TestVmRun(t *testing.T) {

	var vm types.VM = &emscripten.EMVM{}
	vm.Init()
	p := vm.LoadExecFile("/Users/shang/Documents/demo/a.out.wasm")
	vm.InvokeMethod(p, "_main")

}
