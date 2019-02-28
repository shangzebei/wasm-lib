package emscripten

import (
	"github.com/perlin-network/life/exec"
	"wasmgo/wasm"
)

func GlobalCtors(vm *exec.VirtualMachine) {
	wasm.RunMainFunc(vm, "globalCtors")
}
