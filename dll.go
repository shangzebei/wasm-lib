//+build !386

package main

import "C"
import "wasmgo/llvm"

//export load
func load(execFile string) int {
	return llvm.Load(execFile)
}

//export loadExecFile
func loadExecFile(execFile string) int {
	return llvm.LoadExecFile(execFile)
}

//export invokeMethod
func invokeMethod(p int, methodName string) int64 {
	return llvm.InvokeMethod(p, methodName)
}
