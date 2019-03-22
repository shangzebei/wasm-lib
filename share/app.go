package share

import (
	"io/ioutil"
	"log"
	"wasmgo/emscripten"
	"wasmgo/types"
	"wasmgo/wasm"
)

var vm types.VM = &emscripten.EMVM{}

func VmLoad(execFile string) int {
	return vm.Load(execFile)
}

func VmLoadExecFile(execFile string) int {
	return vm.LoadExecFile(execFile)
}

func VmInvokeMethod(p int, methodName string, param []string) int64 {
	return vm.InvokeMethod(p, methodName, param...)
}

func SetVMPlugPath(path string) {
	wasm.SetPlugPath(path)
}

func SetDebug(enable bool) {
	if !enable {
		log.SetOutput(ioutil.Discard)
	}
}
func InitVM() {
	vm.Init()
}
