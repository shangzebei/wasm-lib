package llvm

import (
	"io/ioutil"
	"wasmgo/runtime"
	"wasmgo/wasm"
)

func CallMain(args ...int64) {
	wasm.RegisterFunc(
		&ExportPre{},
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
	)

	input, err := ioutil.ReadFile("/Users/shang/Documents/demo/a.out.wasm")
	if err != nil {
		panic(err)
	}
	wm, _ := wasm.LoadWMFromBytes(input)
	m := VmManger{}
	m.Init(wm, &ExportPre{})
	argc := len(args) + 1
	argv := StackAlloc(wm, (argc+1)*4)
	pos := (argv >> 2) * 4
	copy([]byte("./this.program"), wm.Memory[pos:pos])

	//b, e := json.Marshal(&types.FuncList)
	//fmt.Println(string(b), e)

	wasm.RunFunc(wm, "main")
	//HEAP32[argv >> 2] = allocateUTF8OnStack(Module["thisProgram"]);

}
