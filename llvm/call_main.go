package llvm

import (
	"fmt"
	"io/ioutil"
	"wasmgo/runtime"
	"wasmgo/wasm"
)

func CallMain(wasmFile string, args ...int64) {
	//f, err := os.Open("logfile")
	//if err != nil {
	//}
	//log.SetOutput(f)

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
	)
	input, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		panic(err)
	}
	wm, _ := wasm.LoadWMFromBytes(input)
	m := VmManger{}
	defer func() {
		m.CheckUnflushedContent()
	}()
	m.Init(wm, &VMalloc{wm})
	argc := len(args) + 1
	argv := StackAlloc(wm, (argc+1)*4)
	pos := (argv >> 2) * 4
	copy([]byte("./this.program"), wm.Memory[pos:pos])

	//b, e := json.Marshal(&types.FuncList)
	//fmt.Println(string(b), e)
	fmt.Println(wasm.GetVMemory().Malloc(44))
	fmt.Println(wasm.GetVMemory().Malloc(44))
	fmt.Println(wasm.GetVMemory().Malloc(444))
	fmt.Println(wasm.GetVMemory().Malloc(44))
	//fmt.Println(ZSt18uncaught_exceptionv(wm))

	wasm.RunMainFunc(wm, "main")

	//HEAP32[argv >> 2] = allocateUTF8OnStack(Module["thisProgram"]);

}
