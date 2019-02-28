package llvm

import (
	"github.com/perlin-network/life/exec"
	"io/ioutil"
	"log"
	"wasmgo/runtime"
	"wasmgo/wasm"
)

var moduleList = make([]*exec.VirtualMachine, 0)
var _vm LLVMManger

func init() {
	log.Println("........init........")
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
		//&emscripten.EMSCriptenFun{},
	)
}

func Load(execFile string) int {
	input, err := ioutil.ReadFile(execFile)
	if err != nil {
		panic(err)
	}
	wm, _ := wasm.LoadWMFromBytes(input)
	moduleList = append(moduleList, wm)
	return len(moduleList) - 1
}

func LoadExecFile(execFile string) int {
	_vm = LLVMManger{}
	p := Load(execFile)
	wm := moduleList[p]
	_vm.Init(wm, &VMalloc{Vm: wm})

	//argc := len(args) + 1
	//argv := StackAlloc(wm, (argc+1)*4)
	//pos := (argv >> 2) * 4
	//copy([]byte("./this.program"), wm.Memory[pos:pos])
	return p
}

func InvokeMethod(p int, methodName string, param ...int64) int64 {
	defer func() {
		_vm.CheckUnflushedContent()
	}()
	return wasm.RunMainFunc(moduleList[p], methodName, param...)
}
