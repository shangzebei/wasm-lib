package lib

import (
	"fmt"
	"wasmgo/types"
)

type Log struct {
	types.VMInterface
}

func (l *Log) Log_s(log string) {
	fmt.Println("-----------------------------")
	fmt.Printf("log:: %s\n", log)
	fmt.Println("-----------------------------")
}

func (*Log) Log_i(log int) {
	fmt.Println("-----------------------------")
	fmt.Printf("log:: %d\n", log)
	fmt.Println("-----------------------------")
}

func (*Log) Log_f(log float32) {
	fmt.Println("-----------------------------")
	fmt.Printf("log:: %f\n", log)
	fmt.Println("-----------------------------")
}

func (*Log) Log_l(log int64) {
	fmt.Println("-----------------------------")
	fmt.Printf("log:: %d\n", log)
	fmt.Println("-----------------------------")
}

func (*Log) Log_d(log float64) {
	fmt.Println("-----------------------------")
	fmt.Printf("log:: %f\n", log)
	fmt.Println("-----------------------------")
}
