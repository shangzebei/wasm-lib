package llvm

import (
	"encoding/binary"
	"fmt"
	"github.com/perlin-network/life/exec"
	"math"
	"wasmgo/runtime"
	"wasmgo/types"
	"wasmgo/wasm"
)

type VmManger struct {
	vm             *exec.VirtualMachine //
	only           *types.FuncOnly
	STATIC_BASE    int64 //
	STATICTOP      int64 //
	STACK_BASE     int64 //
	TOTAL_STACK    int64 //
	STACKTOP       int64 //
	STACK_MAX      int64 //
	staticSealed   int64 //
	TOTAL_MEMORY   int64
	GLOBAL_BASE    int64
	DYNAMIC_BASE   int64
	DYNAMICTOP_PTR int64
}

var STACK_ALIGN = int64(16)

func (m *VmManger) Init(vm *exec.VirtualMachine, vMem types.VMemory) {
	wasm.AddFunc(m, "Sbrk")
	wasm.AddFunc(m, "Brk")

	m.vm = vm
	wasm.SetVMemory(vMem)
	m.GLOBAL_BASE = 1024
	m.TOTAL_STACK = HEAP_BASE(vm) - DATA_END(vm)
	m.TOTAL_MEMORY = m.GetTotalMemory()
	m.STATICTOP = m.GLOBAL_BASE + DATA_END(vm)
	m.DYNAMICTOP_PTR = m.StaticAlloc(4)
	m.STACKTOP = m.AlignMemory(m.STATICTOP)
	m.STACK_BASE = m.STACKTOP
	m.STACK_MAX = m.STACK_BASE + m.TOTAL_STACK
	m.DYNAMIC_BASE = m.AlignMemory(m.STACK_MAX)
	binary.LittleEndian.PutUint32(m.vm.Memory[m.DYNAMICTOP_PTR:m.DYNAMICTOP_PTR+4], uint32(m.DYNAMIC_BASE))

	WASM_CALL_CTORS(vm)

}

func (m *VmManger) GetTotalMemory() int64 {
	total := m.vm.Config.DefaultTableSize * m.vm.Config.DefaultMemoryPages
	m.TOTAL_MEMORY = int64(total)
	return m.TOTAL_MEMORY
}

func (m *VmManger) abortOnCannotGrowMemory() bool {
	panic("Cannot enlarge memory arrays. Either (1) compile with  -s TOTAL_MEMORY=X  with X higher than the current value " + ", (2) compile with  -s ALLOW_MEMORY_GROWTH=1  which allows increasing the size at runtime, or (3) if you want malloc to return NULL (0) instead of this abort, compile with  -s ABORTING_MALLOC=0 ")
	return false
}

func (m *VmManger) CheckUnflushedContent() {
	FFLUSH(m.vm, 0)
	s := wasm.GetInstance("SystemCall").(*lib.SystemCall)
	if len(s.Buffers[1].([]byte)) != 0 || len(s.Buffers[2].([]byte)) != 0 {
		fmt.Println("stdio streams had content in them that was not flushed. you should set EXIT_RUNTIME to 1 (see the FAQ), or make sure to emit a newline when you printf etc.")
	}
}

func (m *VmManger) enlargeMemory() bool {
	return m.abortOnCannotGrowMemory()
}

func (m *VmManger) StaticAlloc(size int64) int64 {
	var ret = m.STATICTOP
	m.STATICTOP = m.STATICTOP + size + 15&-16
	return ret
}

func (m *VmManger) AlignMemory(size int64) int64 {
	factor := STACK_ALIGN
	var ret = int64(math.Ceil(float64(size/factor))) * factor
	return ret
}

func (m *VmManger) AllocateUTF8OnStack(str string) int64 {
	size := len([]byte(str)) + 1
	ret := StackAlloc(m.vm, size)
	//stringToUTF8Array(str, HEAP8, ret, size);
	return ret
}

//int sbrk(void *addr);
func (v *VmManger) Sbrk(increment int) int32 {
	increment = increment | 0
	var oldDynamicTop = int32(0)
	var newDynamicTop = int32(0)
	//var totalMemory = int32(0)

	oldDynamicTop = int32(binary.LittleEndian.Uint32(v.vm.Memory[v.DYNAMICTOP_PTR : v.DYNAMICTOP_PTR+4]))
	newDynamicTop = oldDynamicTop + int32(increment)
	//if (increment | 0) > 0 & (newDynamicTop | 0) < (oldDynamicTop | 0) | (newDynamicTop | 0) < 0) {
	//	v.abortOnCannotGrowMemory() | 0;
	//	___setErrNo(12);
	//	return -1
	//}
	//HEAP32[DYNAMICTOP_PTR >> 2] = newDynamicTop
	binary.LittleEndian.PutUint32(v.vm.Memory[v.DYNAMICTOP_PTR:v.DYNAMICTOP_PTR+4], uint32(newDynamicTop))

	totalMemory := int32(v.GetTotalMemory())
	if (newDynamicTop | 0) > (totalMemory | 0) {
		v.enlargeMemory()
		//HEAP32[DYNAMICTOP_PTR>>2] = oldDynamicTop
		//___setErrNo(12);
	}
	return newDynamicTop | 0
}
func (v *VmManger) Brk(newDynamicTop int) int {
	newDynamicTop = newDynamicTop | 0
	var oldDynamicTop = int32(0)
	var totalMemory = int64(0)
	if (newDynamicTop | 0) < 0 {
		v.abortOnCannotGrowMemory()
		//___setErrNo(12)
		return -1
	}
	oldDynamicTop = int32(binary.LittleEndian.Uint32(v.vm.Memory[v.DYNAMICTOP_PTR : v.DYNAMICTOP_PTR+4]))
	binary.LittleEndian.PutUint32(v.vm.Memory[v.DYNAMICTOP_PTR:v.DYNAMICTOP_PTR+4], uint32(newDynamicTop))
	totalMemory = v.GetTotalMemory()
	if int64(newDynamicTop) > totalMemory {
		fmt.Println(oldDynamicTop)
		//if (v.enlargeMemory() | 0) == 0 {
		//	fmt.Println(oldDynamicTop)
		//	//___setErrNo(12);
		//	//HEAP32[DYNAMICTOP_PTR >> 2] = oldDynamicTop;
		//	return -1
		//}
	}
	return 0
}
