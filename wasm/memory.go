package wasm

import (
	"github.com/perlin-network/life/exec"
	"log"
	"wasmgo/types"
)

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

func CopyNewVm(vm *exec.VirtualMachine) *exec.VirtualMachine {
	nvm := exec.VirtualMachine{
		Config:          vm.Config,
		Module:          vm.Module,
		FunctionCode:    vm.FunctionCode,
		FunctionImports: vm.FunctionImports,
		Memory:          vm.Memory,
		CurrentFrame:    -1,
		CallStack:       make([]exec.Frame, exec.DefaultCallStackSize),
		Table:           vm.Table,
		Globals:         vm.Globals,
		//NumValueSlots    int
		//Yielded          int64
		//InsideExecute    bool
		//ExitError        interface{}
		ImportResolver: vm.ImportResolver,
	}
	return &nvm
}
