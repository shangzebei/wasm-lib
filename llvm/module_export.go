package llvm

import (
	"github.com/perlin-network/life/exec"
	"log"
	"wasmgo/wasm"
)

type VMalloc struct {
	Vm *exec.VirtualMachine
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

//(export "dynCall_vi" (func $1530))
//(export "dynCall_iiii" (func $1531))
//(export "dynCall_iii" (func $1532))
//(export "dynCall_v" (func $1533))
//(export "dynCall_ii" (func $1534))
//(export "dynCall_vii" (func $1535))
//(export "dynCall_viiii" (func $1536))
//(export "dynCall_viii" (func $1537))
//(export "dynCall_iiiiii" (func $1538))
//(export "dynCall_iiiiiii" (func $1539))
//(export "dynCall_iiiiid" (func $1540))
//(export "dynCall_iiiiiiiii" (func $1541))
//(export "dynCall_iiiii" (func $1542))
//(export "dynCall_iiiiiiii" (func $1543))
//(export "dynCall_viiiiii" (func $1544))
//(export "dynCall_viiiii" (func $1545))

func DYNCALL_VI(vm *exec.VirtualMachine, p int64, a int64) {
	wasm.RunFunc(vm, "dynCall_vi", p, a)
}

func DYNCALL_IIII(vm *exec.VirtualMachine, p int64, a int64, b int64, c int64) int64 {
	return wasm.RunFunc(vm, "dyncall_iiii", p, a, b, c)
}

func DYNCALL_III(vm *exec.VirtualMachine, p int64, a int64, b int64) int64 {
	return wasm.RunFunc(vm, "dyncall_iii", p, a, b)
}

func DYNCALL_II(vm *exec.VirtualMachine, p int64, a int64) int64 {
	return wasm.RunFunc(vm, "dynCall_ii", p, a)
}

func DYNCALL_V(vm *exec.VirtualMachine, p int64) {
	wasm.RunFunc(vm, "dynCall_v", p)
}

func DYNCALL_VII(vm *exec.VirtualMachine, p int64, a int64, b int64) {
	wasm.RunFunc(vm, "dynCall_vii", p, a, b)
}

func HEAP_BASE(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "__heap_base")
}

func DATA_END(vm *exec.VirtualMachine) int64 {
	return wasm.GetExport(vm, "__data_end")
}

func (v *VMalloc) Malloc(size int64) int64 {
	if v.Vm == nil {
		log.Fatalln("error e.Vm==nil")
	}
	return wasm.RunFunc(v.Vm, "malloc", size)
}

func (v *VMalloc) Free(point int64) int64 {
	if v.Vm == nil {
		log.Fatalln("error e.Vm==nil")
	}
	return wasm.RunFunc(v.Vm, "free", point)
}

func FFLUSH(vm *exec.VirtualMachine, x int) int64 {
	return wasm.RunMainFunc(vm, "fflush", int64(x))
}

func ZSt18uncaught_exceptionv(vm *exec.VirtualMachine) int64 {
	return wasm.RunFunc(vm, "_ZSt18uncaught_exceptionv")
}
