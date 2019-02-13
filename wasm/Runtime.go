package wasm

import (
	"fmt"
	"github.com/perlin-network/life/exec"
	"log"
	"wasmgo/types"
)

type Resolver struct {
	tempRet0 int64
}

// ResolveFunc defines a set of import functions that may be called within a WebAssembly module.
func (r *Resolver) ResolveFunc(module, field string) exec.FunctionImport {
	switch module {
	case "env":
		find := types.FuncList[field]
		if find != nil {
			return InvokeFunc(find)
		} else {
			log.Fatalf("Func unknown field: %s \n", field)
		}
	}
	return func(vm *exec.VirtualMachine) int64 {
		log.Fatalf("Func unknown field: %s 【invoke】 \n", field)
		return 0
	}
}

// ResolveGlobal defines a set of global variables for use within a WebAssembly module.
func (r *Resolver) ResolveGlobal(module, field string) int64 {
	fmt.Printf("Resolve global: %s %s\n", module, field)
	switch module {
	case "env":
		switch field {
		case "__life_magic":
			return 424
		default:
			fmt.Printf("Global unknown field: %s", field)
		}
	default:
		fmt.Printf("Global unknown module: %s", module)
	}
	return 0
}

func LoadWMFromBytes(code []byte) (*exec.VirtualMachine, error) {
	vm, err := exec.NewVirtualMachine(code,
		exec.VMConfig{
			EnableJIT:            false,
			DefaultMemoryPages:   256,
			DefaultTableSize:     65536,
			DisableFloatingPoint: false,
		}, new(Resolver), nil)

	if err != nil {
		panic(err)
	}
	if vm.Module.Base.Start != nil {
		startID := int(vm.Module.Base.Start.Index)
		_, err := vm.Run(startID)
		if err != nil {
			vm.PrintStackTrace()
			panic(err)
		}
	}
	return vm, err
}

func RunFunc(vm *exec.VirtualMachine, name string, params ...int64) int64 {
	entryID, ok := vm.GetFunctionExport(name)
	if !ok {
		fmt.Printf("Entry function %s not found; starting from 0.\n", name)
		entryID = 0
	}

	ret, err := vm.Run(entryID, params...)

	if err != nil {
		vm.PrintStackTrace()
		panic(err)
	}
	return ret
}

func GetExport(vm *exec.VirtualMachine, name string) int64 {
	entryID, ok := vm.GetGlobalExport(name)
	if !ok {
		fmt.Printf("Entry GlobalExport %s not found; starting from 0.\n", name)
		entryID = 0
	}
	return vm.Globals[entryID]
}
