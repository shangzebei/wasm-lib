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
	//e.Replace("AssertFail", "___assert_fail")
}

func (*EMSCriptenFun) NullFuncII(p int32) {
	fmt.Println(p)
}

func (*EMSCriptenFun) AssertFail(a int32, b int32, c int32, d int32) {
	fmt.Println(a, b, c, d)
}

func (em *EMSCriptenFun) GetHeapSize() int32 {
	total := em.Vm.Config.DefaultTableSize * em.Vm.Config.DefaultMemoryPages
	return int32(total)
}
