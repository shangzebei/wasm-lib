//+build !386

package main

import "C"
import (
	"wasmgo/emscripten"
	"wasmgo/types"
	"wasmgo/wasm"
)

var vm types.VM = &emscripten.EMVM{}

//export vmLoad
func vmLoad(execFile string) int {
	return vm.Load(execFile)
}

//export vmLoadExecFile
func vmLoadExecFile(execFile string) int {
	return vm.LoadExecFile(execFile)
}

//export vmInvokeMethod
func vmInvokeMethod(p int, methodName string, param []string) int64 {
	return vm.InvokeMethod(p, methodName, param...)
}

//export setVMPlugPath
func setVMPlugPath(path string) {
	wasm.SetPlugPath(path)
}

//export initVM
func initVM() {
	vm.Init()
}
