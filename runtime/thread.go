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
func (t *Thread) Pthread_mutex_destroy() {

}

func (t *Thread) Pthread_mutex_init() {

}

func (t *Thread) Pthread_mutex_lock(x int) int {
	x = x | 0
	return 0
}

func (t *Thread) Pthread_mutex_trylock(x int) int {
	x = x | 0
	return 0
}

func (t *Thread) Pthread_mutex_unlock(x int) int {
	x = x | 0
	return 0
}
func (t *Thread) Pthread_cond_broadcast(x int) int {
	x = x | 0
	return 0
}

//func ___cxa_atexit() {
// return _atexit.apply(null, arguments)
//}
