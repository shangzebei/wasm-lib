package wasm

import (
	"fmt"
	"github.com/perlin-network/life/exec"
)

//TODO may error
func FindfreeSpece(Vm *exec.VirtualMachine, lenght int64) int64 {
	vmm := Vm.Memory
	//if len(vmm) < int(lenght) {
	//	panic("out of memory in FindfreeSpece")
	//	return 0
	//}
	var pos int64
	var start int64
	for index, value := range vmm {
		if value == 0 {
			if pos == 0 {
				start = int64(index)
			}
			pos++
		} else {
			pos = 0
			start = 0
		}
		if pos > lenght { //100001
			fmt.Printf("find memory pos %d \n", start+1)
			return start + 1
		}
	}
	if pos < lenght {
		fmt.Printf("vmm lenght %d need %d start %d \n ", len(vmm), lenght, start)
		orgLen := int64(len(vmm))
		newLen := int(lenght-(orgLen-start)) + 1
		for i := 0; i <= newLen; i++ {
			Vm.Memory = append(Vm.Memory, 0)
		}
		fmt.Printf("gen memory now %d \n", len(vmm))
	}
	return start + 1
}

func GetTotalMemory(Vm *exec.VirtualMachine) int {
	return len(Vm.Memory)
}

func GetMemoryEndPos(p int64, Vm *exec.VirtualMachine) int64 {
	vmm := Vm.Memory[p:]
	if Vm.Memory[p] == 0 {
		return 0
	}
	for index, value := range vmm {
		if value == 0 {
			return p + int64(index)
		}
	}
	return 0
}
