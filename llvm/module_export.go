package llvm

import (
	"github.com/perlin-network/life/exec"
	"log"
	"wasmgo/wasm"
)

type VMalloc struct {
	vm *exec.VirtualMachine
}

func StackAlloc(vm *exec.VirtualMachine, len int) int64 {
	return wasm.RunMainFunc(vm, "stackAlloc", int64(len))
}

func TOTAL_STACK(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "TOTAL_STACK")
}

func TOTAL_MEMORY(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "TOTAL_MEMORY")
}

func WASM_CALL_CTORS(vm *exec.VirtualMachine) {
	wasm.RunMainFunc(vm, "__wasm_call_ctors")
}

func DYNCALL_IIII(vm *exec.VirtualMachine, params ...int64) {
	//wasm.RunFunc(vm, "dyncall_iiii", params)
}

func HEAP_BASE(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "__heap_base")
}

func DATA_END(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "__data_end")
}

func (v *VMalloc) Malloc(size int64) int64 {
	if v.vm == nil {
		log.Fatalln("error e.Vm==nil")
	}
	return wasm.RunFunc(v.vm, "malloc", size)
}

func (v *VMalloc) Free(point int64) int64 {
	if v.vm == nil {
		log.Fatalln("error e.Vm==nil")
	}
	return wasm.RunFunc(v.vm, "free", point)
}

func FFLUSH(vm *exec.VirtualMachine, x int) int64 {
	return wasm.RunMainFunc(vm, "fflush", int64(x))
}

func ZSt18uncaught_exceptionv(vm *exec.VirtualMachine) int64 {
	return wasm.RunFunc(vm, "_ZSt18uncaught_exceptionv")
}
