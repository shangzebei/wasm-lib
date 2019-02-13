package lib

import (
	"fmt"
	"wasmgo/types"
)

type Thread struct {
	types.RegInterface
}

func (t *Thread) Init() {
	t.Replace("Lock", "__lock")
	t.Replace("UnLock", "__unlock")
}

func (t *Thread) Lock(a int) {
	fmt.Println("no Lock")
}
func (t *Thread) UnLock(a int) {
	fmt.Println("no UnLock")
}
