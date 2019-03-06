//+build !386

package main

import "C"
import (
	"wasmgo/emscripten"
	"wasmgo/types"
	"wasmgo/wasm"
)

var vm types.VM = &emscripten.EMVM{}

//export load
func load(execFile string) int {
	return vm.Load(execFile)
}

//export loadExecFile
func loadExecFile(execFile string) int {
	return vm.LoadExecFile(execFile)
}

//export invokeMethod
func invokeMethod(p int, methodName string) int64 {
	return vm.InvokeMethod(p, methodName)
}

//export setPlugPath
func setPlugPath(path string) {
	wasm.SetPlugPath(path)
}

//export initVM
func initVM() {
	vm.Init()
}
