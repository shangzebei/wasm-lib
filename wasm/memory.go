package wasm

import (
	"github.com/perlin-network/life/exec"
	"log"
	"wasmgo/types"
)

func GetTotalMemory(Vm *exec.VirtualMachine) int {
	return len(Vm.Memory)
}

var _vMem types.VMemory

func GetVMemory() types.VMemory {
	if _vMem == nil {
		log.Fatalf("you must impl and SetVMemory ...")
	}
	return _vMem
}
func SetVMemory(v types.VMemory) {
	_vMem = v
}
