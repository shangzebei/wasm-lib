package emscripten

import (
	"encoding/json"
	"fmt"
	"github.com/perlin-network/life/exec"
	"io/ioutil"
	"wasmgo/runtime"
	"wasmgo/types"
	"wasmgo/util"
	"wasmgo/wasm"
)

type EMVM struct {
}

var moduleList = make([]*exec.VirtualMachine, 0)
var _vm EMscriptenManger

func (emvm *EMVM) Load(execFile string) int {
	if !util.Exists(execFile) {
		fmt.Println("file not exists! ")
		return 0
	}
	input, err := ioutil.ReadFile(execFile)
	if err != nil {
		panic(err)
	}
	wm, _ := wasm.LoadWMFromBytes(input)
	moduleList = append(moduleList, wm)
	return len(moduleList) - 1
}

func (emvm *EMVM) LoadExecFile(execFile string) int {
	_vm = EMscriptenManger{}
	var p int
	_vm.Init(func() *exec.VirtualMachine {
		p = emvm.Load(execFile)
		wm := moduleList[p]
		return wm
	})
	//argc := len(args) + 1
	//argv := StackAlloc(wm, (argc+1)*4)
	//pos := (argv >> 2) * 4
	//copy([]byte("./this.program"), wm.Memory[pos:pos])
	return p
}

func (emvm *EMVM) InvokeMethod(p int, methodName string, param ...string) int64 {
	defer func() {
		//_vm.CheckUnflushedContent()
	}()
	return wasm.RunMainFunc(moduleList[p], methodName, util.DisposeParam(param, moduleList[p])...)
}

func (emvm *EMVM) Init() {
	wasm.RegisterFunc(
		&lib.Exception{},
		&lib.Log{},
		&lib.Math{},
		&lib.MemoryInterface{},
		&lib.String{},
		&lib.StdLib{RegInterface: types.RegInterface{ReplaceSymbol: map[string]string{"__buildEnvironment": "___buildEnvironment"}}}, //__buildEnvironment
		&lib.Encrypt{},
		&lib.Time{},
		&lib.Http{RegInterface: types.RegInterface{ReplaceSymbol: map[string]string{"Http": "_Http"}}},
		&lib.SystemCall{RegInterface: types.RegInterface{ReplaceSymbol: map[string]string{"__": "___"}}},
		&lib.Thread{},
		&lib.System{},

		&EMSCriptenFun{},
	)

	b, _ := json.Marshal(&types.FuncList)
	fmt.Println(string(b))

}
