package wasm

import (
	"log"
	"wasmgo/types"
)

var _vMem types.VMemory

func GetVMemory() types.VMemory {
	if _vMem == nil {
		log.Fatalf("you must impl and SetVMemory ...")
	}
	return _vMem
}
func SetVMemory(v types.VMemory) {
	_vMem = v
}
