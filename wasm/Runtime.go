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
	switch module {
	case "env":
		if types.GlobalList[field] != nil {
			return types.GlobalList[field].(int64)
		} else {
			log.Printf("env unknown field: %s\n", field)
		}
	case "global":
		switch field {
		default:
			log.Printf("Global unknown field: %s\n", field)
		}

	default:
		fmt.Printf("Global unknown module: %s\n", module)
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

func RunMainFunc(vm *exec.VirtualMachine, name string, params ...int64) int64 {
	entryID, ok := vm.GetFunctionExport(name)
	if !ok {
		fmt.Printf("Entry function %s not found\n", name)
		return 0
	}
	ret, err := vm.Run(entryID, params...)

	if err != nil {
		//vm.PrintStackTrace()
		panic(err)
	}
	return ret
}

func RunFunc(vm *exec.VirtualMachine, name string, params ...int64) int64 {
	entryID, ok := vm.GetFunctionExport(name)
	aa := CopyNewVm(vm)
	if !ok {
		fmt.Printf("Entry function %s not found; starting from 0.\n", name)
		entryID = 0
	}
	ret, err := aa.Run(entryID, params...)

	if err != nil {
		//vm.PrintStackTrace()
		panic(err)
	}
	return ret
}

/**
 * invoke point method
 */
func InvokeMethod(vm *exec.VirtualMachine, functionId int, param ...int64) int64 {
	ret, err := CopyNewVm(vm).Run(int(vm.Table[functionId]), param...)
	if err != nil {
		vm.PrintStackTrace()
		panic(err)
	}
	return ret
}

func GetExport(vm *exec.VirtualMachine, name string) (int64, bool) {
	entryID, ok := vm.GetGlobalExport(name)
	if !ok {
		fmt.Printf("Entry GlobalExport %s not found \n", name)
		entryID = 0
	}
	return vm.Globals[entryID], ok
}
