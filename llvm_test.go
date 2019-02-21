package main

import (
	"testing"
	"wasmgo/llvm"
)

func TestVmRun(t *testing.T) {
	//log.SetOutput(ioutil.Discard)
	p := llvm.LoadExecFile("/Users/shang/Documents/demo/a.out.wasm")
	llvm.InvokeMethod(p, "main")

}
