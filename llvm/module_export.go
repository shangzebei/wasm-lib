package llvm

import (
	"github.com/perlin-network/life/exec"
	"wasmgo/types"
	"wasmgo/wasm"
)

type ExportPre struct {
	types.PreFuncInf
	types.VMInterface
}

func StackAlloc(vm *exec.VirtualMachine, len int) int64 {
	return wasm.RunFunc(vm, "stackAlloc", int64(len))
}

func TOTAL_STACK(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "TOTAL_STACK")
}

func TOTAL_MEMORY(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "TOTAL_MEMORY")
}

func WASM_CALL_CTORS(vm *exec.VirtualMachine) {
	wasm.RunFunc(vm, "__wasm_call_ctors")
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

func (e *ExportPre) Malloc(size int64) int64 {
	return wasm.RunFunc(e.Vm, "MALLOC")
}

func (e *ExportPre) Free(point int64) int64 {
	return wasm.RunFunc(e.Vm, "free")
}

func FFLUSH(vm *exec.VirtualMachine, x int) int64 {
	return wasm.RunFunc(vm, "fflush", int64(x))
}
