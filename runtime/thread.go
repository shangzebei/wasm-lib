package lib

import (
	"sync"
	"wasmgo/types"
	"wasmgo/wasm"
)

type Thread struct {
	types.RegInterface
	m sync.Mutex
	types.VMInterface
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
func (t *Thread) Pthread_create(thread int64, attr int64, start_routine int, arg int64) int64 {
	wasm.InvokeMethod(t.Vm, start_routine, 1)
	return 0
}

//pthread_exit (status)
func (t *Thread) Pthread_exit(status int64) {

}

//int pthread_detach(pthread_t);
func (t *Thread) Pthread_detach(threadId int64) int64 {
	return 0
}

//int pthread_join(pthread_t , void * _Nullable * _Nullable)
func (t *Thread) Pthread_join(threadId int64, p int64) int64 {
	return 0
}

//int pthread_key_create(pthread_key_t *, void (* _Nullable)(void *));
func (t *Thread) Pthread_key_create(pthread_key_t int64, p int64) int64 {
	//fmt.Println(int(C.random()))
	return 0
}

//int pthread_setspecific(pthread_key_t , const void * _Nullable);
func (t *Thread) Pthread_setspecific(pthread_key_t int64, p int64) int64 {
	return -1
}
