package main

import (
	"testing"
)

func TestVmRun(t *testing.T) {
	//log.SetOutput(ioutil.Discard)
	p := LoadExecFile("/Users/shang/Documents/demo/a.out.wasm")
	InvokeMethod(p, "main")

}
