package util

import (
	"github.com/perlin-network/life/exec"
	"os"
	"strconv"
	"wasmgo/wasm"
)

func AllocChars(strings string, vm *exec.VirtualMachine) int64 {
	return AllocBytes([]byte(strings), vm)
}

func AllocBytes(bytes []byte, vm *exec.VirtualMachine) int64 {
	l := len(bytes)
	p := wasm.GetVMemory().Malloc(int64(l))
	copy(vm.Memory[p:int(p)+l], bytes)
	return p
}

/**
 *
 */
func CheckIFElse(condition int64, defalt int64) int64 {
	if condition <= 0 {
		return defalt
	} else {
		return condition
	}
}

func DisposeParam(args []string, machine *exec.VirtualMachine) []int64 {
	var re = make([]int64, len(args))
	for index, value := range args {
		v, e := strconv.ParseInt(value, 10, 64)
		if e != nil {
			re[index] = AllocChars(value, machine)
		} else {
			re[index] = v
		}
	}
	return re
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
