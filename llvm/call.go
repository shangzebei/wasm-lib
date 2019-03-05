package llvm

import (
	"github.com/perlin-network/life/exec"
	"io/ioutil"
	"wasmgo/runtime"
	"wasmgo/util"
	"wasmgo/wasm"
)

type LLVM struct {
}

var moduleList = make([]*exec.VirtualMachine, 0)
var _vm LLVMManger

func (llvm *LLVM) Load(execFile string) int {
	input, err := ioutil.ReadFile(execFile)
	if err != nil {
		panic(err)
	}
	wm, _ := wasm.LoadWMFromBytes(input)
	moduleList = append(moduleList, wm)
	return len(moduleList) - 1
}

func (llvm *LLVM) LoadExecFile(execFile string) int {
	_vm = LLVMManger{}
	p := llvm.Load(execFile)
	wm := moduleList[p]
	_vm.Init(wm, &VMalloc{Vm: wm})

	//argc := len(args) + 1
	//argv := StackAlloc(wm, (argc+1)*4)
	//pos := (argv >> 2) * 4
	//copy([]byte("./this.program"), wm.Memory[pos:pos])
	return p
}

func (llvm *LLVM) InvokeMethod(p int, methodName string, param ...string) int64 {
	defer func() {
		_vm.CheckUnflushedContent()
	}()

	return wasm.RunMainFunc(moduleList[p], methodName, util.DisposeParam(param, moduleList[p])...)
}

func (llvm *LLVM) Init() {
	wasm.RegisterFunc(
		&lib.Exception{},
		&lib.Log{},
		&lib.Math{},
		&lib.MemoryInterface{},
		&lib.String{},
		&lib.StdLib{},
		&lib.Encrypt{},
		&lib.Time{},
		&lib.Http{},
		&lib.SystemCall{},
		&lib.Thread{},
		&lib.System{},
	)
}
