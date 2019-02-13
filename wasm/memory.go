package wasm

import (
	"github.com/perlin-network/life/exec"
)

func GetTotalMemory(Vm *exec.VirtualMachine) int {
	return len(Vm.Memory)
}
