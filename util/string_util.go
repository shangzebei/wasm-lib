package util

import (
	"github.com/perlin-network/life/exec"
	"wasmgo/wasm"
)

func AllocChars(strings string, vm *exec.VirtualMachine) int64 {
	return AllocBytes([]byte(strings), vm)
}

func AllocBytes(bytes []byte, vm *exec.VirtualMachine) int64 {
	l := len(bytes)
	p := wasm.GetVMemory().Malloc(int64(l))
	copy(vm.Memory[p:int(p)+l], bytes)
	return p
}

/**
 *
 */
func CheckIFElse(condition int64, defalt int64) int64 {
	if condition <= 0 {
		return defalt
	} else {
		return condition
	}
}
