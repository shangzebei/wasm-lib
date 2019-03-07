package wasm

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/perlin-network/life/exec"
	"log"
	"math"
	"reflect"
	"strings"
	"wasmgo/types"
)

/**
 *
 */
func RegisterFunc(ins ...interface{}) {
	for _, value := range ins {
		types.InstanceList[strings.ToLower(reflect.TypeOf(value).Elem().Name())] = value
		if !IsInterface(value, types.PreFuncInf{}) {
			RegFunc(value)
		}
	}
	proLoad := GetInterfaceByType(ins, types.PreFuncInf{})
	for _, value := range proLoad {
		RegFunc(value)
	}
}

func GetInterfaceByType(ins []interface{}, ty interface{}) []interface{} {
	var funs = make([]interface{}, 0)
	id := getStructId(reflect.TypeOf(ty))
	for _, value := range ins {
		_, ok := reflect.TypeOf(value).Elem().FieldByName(id)
		if ok {
			funs = append(funs, value)
		}
	}
	return funs
}

func IsInterface(org interface{}, inf interface{}) bool {
	id := getStructId(reflect.TypeOf(inf))
	typ := reflect.TypeOf(org)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	_, ok := typ.FieldByName(id)
	return ok
}

func getStructId(typ reflect.Type) string {
	var st = typ
	if typ.Kind() == reflect.Ptr {
		st = typ.Elem()
	}
	return st.Name() + "_" + getMD5(st.Name())
}
func getMD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func RegFunc(ins interface{}) {
	fooType := reflect.TypeOf(ins)
	// auto invoke init method

	init, _ := fooType.MethodByName("Init")

	if init.Func.IsValid() {
		init.Func.Call([]reflect.Value{reflect.ValueOf(ins)})
	}

	if IsInterface(ins, types.FuncOnly{}) {
		return
	}
	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		if method.Name == "Get" || method.Name == "Init" || method.Name == "Replace" {
			continue
		}
		m := GetMethodMeta(ins, method)
		orgMethodName := m.MethodName
		if IsInterface(ins, types.RegInterface{}) { //reg
			get, _ := fooType.MethodByName("Get")
			rem := get.Func.Call([]reflect.Value{reflect.ValueOf(ins), reflect.ValueOf(m.MethodName)})
			replaceSymbol := reflect.ValueOf(ins).Elem().FieldByName("ReplaceSymbol").Interface().(map[string]string)
			if len(replaceSymbol) != 0 {
				for key, value := range replaceSymbol {
					if len(rem) != 0 && rem[0].String() != "" {
						orgMethodName = strings.ReplaceAll(rem[0].String(), key, value)
					} else {
						orgMethodName = strings.ReplaceAll(orgMethodName, key, value)
					}
				}
			} else {
				if len(rem) != 0 && rem[0].String() != "" {
					orgMethodName = rem[0].String()
				}
			}

		}
		types.FuncList[FirstCharLower(orgMethodName)] = m

	}

}

/**
 * if has overload
 */
func AddFunc(ins interface{}, funName string) {
	method, find := reflect.TypeOf(ins).MethodByName(funName)
	if find {
		m := GetMethodMeta(ins, method)
		types.FuncList[FirstCharLower(m.MethodName)] = m
	}
}

func GetMethodMeta(ins interface{}, method reflect.Method) *types.MethodType {
	var med []string
	var ret []string
	//get in type
	for f := 0; f < method.Type.NumIn(); f++ {
		med = append(med, method.Type.In(f).String())
	}
	//get out type
	for k := 0; k < method.Type.NumOut(); k++ {
		ret = append(ret, method.Type.Out(k).String())
	}
	return &types.MethodType{
		ins,
		method.Name,
		med,
		ret,
	}
}

//support overload
func GetFuncName(s string) string {
	if strings.Contains(s, "_") {
		ss := strings.Split(s, "_")
		return FirstCharLower(ss[0])
	} else {
		return FirstCharLower(s)
	}
}

func GetInstance(name string) interface{} {
	return types.InstanceList[strings.ToLower(name)]
}

func InvokeFunc(methInfo interface{}) func(vm *exec.VirtualMachine) int64 {
	metype := methInfo.(*types.MethodType)
	return func(vm *exec.VirtualMachine) int64 {
		frame := vm.GetCurrentFrame()
		//deal param
		var param = make([]reflect.Value, 10)
		for a := 0; a < len(metype.Types); a++ {
			wasmParamIndex := a - 1
			switch metype.Types[a] {
			case types.INT:
				param[a] = reflect.ValueOf(int(frame.Locals[wasmParamIndex]))
			case types.INT64:
				param[a] = reflect.ValueOf(frame.Locals[wasmParamIndex])
			case types.INT32:
				param[a] = reflect.ValueOf(int32(frame.Locals[wasmParamIndex]))
			case types.STRING:
				ptr := frame.Locals[wasmParamIndex]
				stringPtr := GetString(ptr, vm)
				param[a] = reflect.ValueOf(stringPtr)
			case types.FLOAT64:
				param[a] = reflect.ValueOf(math.Float64frombits(uint64(frame.Locals[wasmParamIndex])))
			case types.FLOAT32:
				param[a] = reflect.ValueOf(math.Float32frombits(uint32(frame.Locals[wasmParamIndex])))
			default:
				if a == 0 {
					//GetInstance
					v := reflect.ValueOf(metype.This)
					param[a] = v
					vk := v.Elem().FieldByName("Vm")
					if vk.CanSet() {
						vk.Set(reflect.ValueOf(vm))
					}
				} else {
					println("not support type" + metype.Types[a])
				}
			}
		}
		//invoke function
		method, _ := reflect.TypeOf(metype.This).MethodByName(metype.MethodName)
		log.Printf("call %s \n", method.Name)
		ret := method.Func.Call(param[0:len(metype.Types)])
		if metype.ReturnTypes == nil {
			return 0
		}
		switch metype.ReturnTypes[0] {
		case types.INT, types.INT32, types.INT64:
			return ret[0].Int()
		case types.STRING:
			return SetString(ret[0].String(), vm)
		case types.FLOAT32, types.FLOAT64:
			return int64(math.Float64bits(ret[0].Float()))
		}
		return 0
	}
}

func GetString(ptr int64, vm *exec.VirtualMachine) string {
	if ptr == 0 {
		panic("nil point")
		return ""
	}
	msg := vm.Memory[ptr:]
	for index, value := range msg {
		if value == 0 {
			s := string(vm.Memory[ptr : ptr+int64(index)])
			return s
		}
	}
	return string(vm.Memory[ptr:])
}

func SetString(s string, vm *exec.VirtualMachine) int64 {
	msg := vm.Memory
	p := GetVMemory().Malloc(int64(len(s)))
	copy(msg[p:], []byte(s))
	defer func() {
		GetVMemory().Free(p)
	}()
	return p
}

func FirstCharLower(s string) string {

	if "" == s {
		log.Fatalf("err s %s", s)
	}
	index := strings.LastIndex(s, "_")
	if index == -1 {
		return strings.ToLower(s[0:1]) + s[1:]
	} else {
		return strings.ToLower(s[0:index]) + s[index:]
	}

}

func FirstCharUpper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
