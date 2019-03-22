package emscripten

import (
	"fmt"
	"wasmgo/types"
)

type EMSCriptenFun struct {
	types.RegInterface
	types.VMInterface
}

func (e *EMSCriptenFun) Init() {
	e.Replace("NullFuncII", "nullFunc_ii")
	e.Replace("GetHeapSize", "_emscripten_get_heap_size")
	e.Replace("ResizeHeap", "_emscripten_resize_heap")
	//e.Replace("AssertFail", "___assert_fail")
}

func (*EMSCriptenFun) NullFuncII(p int32) {
	fmt.Println(p)
}

func (*EMSCriptenFun) AssertFail(a int32, b int32, c int32, d int32) {
	fmt.Println(a, b, c, d)
}

func (em *EMSCriptenFun) GetHeapSize() int32 {
	return int32(len(em.Vm.Memory))
}

func (em *EMSCriptenFun) ResizeHeap(requestedSize int32) int32 {
	em.Vm.Memory = append(em.Vm.Memory, make([]byte, requestedSize)...)
	return requestedSize
}
