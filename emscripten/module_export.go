package emscripten

import (
	"github.com/perlin-network/life/exec"
	"wasmgo/wasm"
)

func GlobalCtors(vm *exec.VirtualMachine) {
	wasm.RunMainFunc(vm, "globalCtors")
}

func EstablishStackSpace(vm *exec.VirtualMachine, STACK_BASE int64, STACK_MAX int64) {
	wasm.RunMainFunc(vm, "establishStackSpace", STACK_BASE, STACK_MAX) //establishStackSpace
}

//get STACK_BASE
func StackSave(vm *exec.VirtualMachine) int64 {
	return wasm.RunMainFunc(vm, "stackSave")
}
