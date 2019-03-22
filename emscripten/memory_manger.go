package emscripten

import (
	"github.com/perlin-network/life/exec"
	"log"
	"wasmgo/types"
	"wasmgo/util"
	"wasmgo/wasm"
)

type EMscriptenManger struct {
	vm             *exec.VirtualMachine //
	STATIC_BASE    int64
	STACK_BASE     int64
	STACKTOP       int64
	STACK_MAX      int64
	DYNAMIC_BASE   int64
	DYNAMICTOP_PTR int64
	TOTAL_STACK    int64
	TOTAL_MEMORY   int64
}

type VMImpl struct {
	Vm *exec.VirtualMachine
}

func (em *EMscriptenManger) Init(f func() *exec.VirtualMachine) {

	em.STATIC_BASE = 1024
	em.STACK_BASE = 39120
	em.STACKTOP = em.STACK_BASE
	em.STACK_MAX = 5282000
	em.DYNAMIC_BASE = 5282000
	em.DYNAMICTOP_PTR = 38864
	em.TOTAL_STACK = 5242880

	//assert(STACK_BASE % 16 === 0, 'stack must start aligned');
	//assert(DYNAMIC_BASE % 16 === 0, 'heap must start aligned');
	types.GlobalList["__memory_base"] = int64(1024) // tell the memory segments where to place themselves
	types.GlobalList["__table_base"] = int64(0)     // table starts at 0 by default (even in dynamic linking, for the main module)
	types.GlobalList["tempDoublePtr"] = int64(22560)
	types.GlobalList["DYNAMICTOP_PTR"] = em.DYNAMICTOP_PTR

	em.vm = f()

	wasm.SetVMemory(&VMImpl{Vm: em.vm})

	util.Put32(em.vm.Memory, 0, 0x63736d65)

	util.Put32(em.vm.Memory, em.DYNAMICTOP_PTR, uint32(em.STACK_BASE))

	GlobalCtors(em.vm)

}

func (m *EMscriptenManger) GetTotalMemory() int64 {
	total := m.vm.Config.DefaultTableSize * m.vm.Config.DefaultMemoryPages
	m.TOTAL_MEMORY = int64(total)
	return m.TOTAL_MEMORY
}

func (v *VMImpl) Malloc(size int64) int64 {
	if v.Vm == nil {
		log.Fatalln("error e.Vm==nil")
	}
	return wasm.RunFunc(v.Vm, "_malloc", size)
}

func (v *VMImpl) Free(point int64) int64 {
	if v.Vm == nil {
		log.Fatalln("error e.Vm==nil")
	}
	return wasm.RunFunc(v.Vm, "_free", point)
}
