package types

import (
	"github.com/perlin-network/life/exec"
)

/**
 *    GO    C
 *--------------------
 *  int64   long
 *  float64 double
 *  float32 float
 *  int     int
 *  int32   unsigned long
 *  int64   long int
 */

const (
	STRING  = "string"
	INT     = "int"
	INT32   = "int32"
	INT64   = "int64"
	FLOAT64 = "float64"
	FLOAT32 = "float32"
)

type MethodType struct {
	This        interface{} `json:"-"`
	MethodName  string      `json:"methodName"`
	Types       []string    `json:"types"`
	ReturnTypes []string    `json:"returnTypes"`
}

/**
 * this use for vm
 */
type VMInterface struct {
	VMInterface_5239b67420ae6dda8929ad26f5b66219 int
	Vm                                           *exec.VirtualMachine
}

/**
 * this for replace fun name
 * if struct has Init, only replace method name
 */
type RegInterface struct {
	RegInterface_de0f4ef1220c860d1a3708f94a5a7da1 int
	RePMap                                        map[string]string
}

/**
 * if impl invoke first
 */
type PreFuncInf struct {
	PreFuncInf_3ca04d88600e1136dd1a21a51edf070b int
}

/**
 * not scan all method
 */
type FuncOnly struct {
	FuncOnly_0ce668f932cf805f5821d2ec2cb39403 int
}

/**
 * may Impl
 */
type VMemory interface {
	Malloc(size int64) int64 //this method alloc
	Free(point int64) int64  //this method free
}

/**
 * get the real func,by the runtime
 */
func (r *RegInterface) Get(name string) string {
	return r.RePMap[name]
}

func (r *RegInterface) Replace(a string, b string) {
	if r.RePMap == nil {
		r.RePMap = make(map[string]string)
	}
	r.RePMap[a] = b
}

/**
 * glob funList
 */
var FuncList = make(map[string]interface{})
var GlobalList = make(map[string]interface{})
var InstanceList = make(map[string]interface{})
var PreFuncList = make(map[string]interface{})
