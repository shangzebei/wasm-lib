package lib

import (
	"sync"
	"wasmgo/types"
)

type Thread struct {
	types.RegInterface
	m sync.Mutex
}

func (t *Thread) Init() {
	t.Replace("Lock", "__lock")
	t.Replace("UnLock", "__unlock")
}

func (t *Thread) Lock(a int) {
	t.m.Lock()
}

func (t *Thread) UnLock(a int) {
	t.m.Unlock()
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

//pthread_create (thread, attr, start_routine, arg)
func (t *Thread) Pthread_create(thread int64, attr int64, start_routine int64, arg int64) int64 {
	return -1
}

//pthread_exit (status)
func (t *Thread) Pthread_exit(status int64) {

}
