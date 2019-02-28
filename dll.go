//+build !386

package main

import "C"
import (
	"wasmgo/types"
)

var vm types.VM

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
