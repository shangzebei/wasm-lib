package emscripten

import (
	"github.com/perlin-network/life/exec"
	"wasmgo/llvm"
	"wasmgo/types"
	"wasmgo/wasm"
)

type EMscriptenManger struct {
	STATIC_BASE    int64
	STACK_BASE     int64
	STACKTOP       int64
	STACK_MAX      int64
	DYNAMIC_BASE   int64
	DYNAMICTOP_PTR int64
	TOTAL_STACK    int64
}

func (em *EMscriptenManger) Init(f func() *exec.VirtualMachine) {

	em.STATIC_BASE = 1024
	em.STACK_BASE = 22576
	em.STACKTOP = em.STACK_BASE
	em.STACK_MAX = 5265456
	em.DYNAMIC_BASE = 5265456
	em.DYNAMICTOP_PTR = 22320
	em.TOTAL_STACK = 5242880

	//assert(STACK_BASE % 16 === 0, 'stack must start aligned');
	//assert(DYNAMIC_BASE % 16 === 0, 'heap must start aligned');
	types.GlobalList["__memory_base"] = int64(1024) // tell the memory segments where to place themselves
	types.GlobalList["__table_base"] = int64(0)     // table starts at 0 by default (even in dynamic linking, for the main module)
	types.GlobalList["tempDoublePtr"] = int64(22560)
	types.GlobalList["DYNAMICTOP_PTR"] = em.DYNAMICTOP_PTR

	wasm.SetVMemory(&llvm.VMalloc{})
	vm := f()

	GlobalCtors(vm)

}
