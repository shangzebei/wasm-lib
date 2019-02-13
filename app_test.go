package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"testing"
	"time"
	"wasmgo/llvm"
	"wasmgo/runtime"
	"wasmgo/types"
	"wasmgo/wasm"
)

func TestWasmRun(t *testing.T) {

	wasm.RegisterFunc(
		&llvm.ExportPre{},
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

	b, e := json.Marshal(&types.FuncList)
	fmt.Println(string(b), e)

	input, err := ioutil.ReadFile("/Users/shang/Documents/demo/a.out.wasm")
	if err != nil {
		panic(err)
	}
	wm, _ := wasm.LoadWMFromBytes(input)

	fmt.Println(len(wm.Memory))
	wasm.RunFunc(wm, "main")

}

func TestType(t *testing.T) {
	f := int64(math.Float32bits(89.76))
	fmt.Println(f)
	fmt.Println(math.Float32frombits(uint32(1119061279)))
	fmt.Println(reflect.ValueOf(float32(87.98)).Interface())

}

type AAA struct {
}

func (*AAA) Show(aaa float32) {
	fmt.Println(aaa)
}
func TestFloat(t *testing.T) {
	a := &AAA{}
	me, _ := reflect.TypeOf(a).MethodByName("Show")
	f := math.Float32frombits(uint32(int64(1119061279)))
	fmt.Println(f)
	me.Func.Call([]reflect.Value{
		reflect.ValueOf(a),
		reflect.ValueOf(f),
	})

}
func TestTime(t *testing.T) {
	fmt.Println(time.Now().Unix())
}

func TestF64toi32(t *testing.T) {
	//a:=int64(18328)
	//a:=uint32(18328.1)
	a := float32(18328.12)
	b := int64(math.Float64bits(float64(int32(a))))
	fmt.Println(b)
	fmt.Println(math.Float64frombits(uint64(b)))

}
func TestInt(t *testing.T) {

}

func TestTypeR(t *testing.T) {
	b := lib.SystemCall{}
	m, _ := reflect.TypeOf(b).MethodByName("__syscall6")
	fmt.Println(m.Func.CanSet())
}
