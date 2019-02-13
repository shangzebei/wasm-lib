package lib

import (
	"fmt"
	"wasmgo/types"
	"wasmgo/wasm"
)

type String struct {
	types.VMInterface
}

//TODO
func (s *String) Strcat(dest int64, src int64) int64 {
	end := wasm.GetMemoryEndPos(dest, s.Vm)
	fmt.Println(dest, src)
	fmt.Println(wasm.GetString(dest, s.Vm), wasm.GetString(src, s.Vm))
	fmt.Println(string(s.Vm.Memory))
	fmt.Println(end)
	return 0
}
func (s *String) Strlen(str int64) int64 {
	return wasm.GetMemoryEndPos(str, s.Vm) - str
}
