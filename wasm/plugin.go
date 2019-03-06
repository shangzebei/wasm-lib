package wasm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"reflect"
	"wasmgo/types"
)

var ins = &plugs{make(map[string]*PlugFactory)}

type PlugFactory struct {
	funcPl map[string]*plugin.Plugin
}

type plugs struct {
	Pls map[string]*PlugFactory
}

const WASM_LIB = "/Users/shang/go/src/wasmgo/"

func init() {
	env := os.Getenv("WASM_LIB")
	if env == "" {
		env = WASM_LIB
	}
	log.Printf("WASM_LIB = %s \n", env)
	loadSystem(env)
}

func PlugInstants(plugName string) *PlugFactory {
	v, e := ins.Pls[plugName]
	if !e {
		return &PlugFactory{}
	}
	return v
}

func LoadPlugin(file string) {
	p, err := plugin.Open(file)
	if err != nil {
		fmt.Printf("eror load plugin %s \n", file)
		return
	}
	init, err := p.Lookup("Init")
	if err == nil {
		vmPlug := &types.VMPlugin{}
		init.(func(vmPlugin *types.VMPlugin))(vmPlug)
		if vmPlug.PlugName == "" {
			return
		}
		regs := vmPlug.GetRegs()
		ins.Pls[vmPlug.PlugName] = &PlugFactory{make(map[string]*plugin.Plugin)}
		for _, value := range regs {
			ins.Pls[vmPlug.PlugName].funcPl[value] = p
		}
		log.Printf("load path %s success ! Name= %s \n", file, vmPlug.PlugName)
	}

}

func loadSystem(env string) {
	log.Println("load system lib")
	files, _ := filepath.Glob(env + "*.so")
	for _, value := range files {
		LoadPlugin(value)
	}
}

func (plu *PlugFactory) Call(methodName string, param ...interface{}) []interface{} {

	sym, has := plu.GetMethod(methodName)
	if has {
		var vParam []reflect.Value
		for _, value := range param {
			vParam = append(vParam, reflect.ValueOf(value))
		}
		v := reflect.ValueOf(sym).Call(vParam)
		var rParam []interface{}
		for _, value := range v {
			rParam = append(rParam, value.Interface())
		}
		return rParam
	}
	return nil

}

func (plu *PlugFactory) GetMethod(methodName string) (interface{}, bool) {
	if plu.funcPl == nil {
		return nil, false
	}
	if plu.funcPl[methodName] == nil {
		fmt.Printf("no such method %s\n", methodName)
		return nil, false
	} else {
		m, _ := plu.funcPl[methodName].Lookup(methodName)
		return m, true
	}
}
