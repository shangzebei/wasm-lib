package lib

import (
	"log"
	"wasmgo/types"
	"wasmgo/wasm"
)

type Exception struct {
	types.RegInterface
}

func (e *Exception) Init() {
	e.Replace("Cxa_allocate_exception", "__cxa_allocate_exception")
	e.Replace("Cxa_throw", "__cxa_throw")
	e.Replace("Cxa_uncaught_exception", "__cxa_uncaught_exception")
}

func (e *Exception) Cxa_allocate_exception(size int64) int64 {
	return wasm.GetVMemory().Malloc(size)
}
func (e *Exception) Cxa_throw(ptr int64, typ int64, destructor int64) {
	log.Fatalf("ptr= %d - Exception catching is disabled, this exception cannot be caught. Compile with -s DISABLE_EXCEPTION_CATCHING=0 or DISABLE_EXCEPTION_CATCHING=2 to catch.", ptr)
}

func (e *Exception) Cxa_uncaught_exception() int32 {
	return 0
}
func (*Exception) AbortStackOverflow(p int32) {

}
