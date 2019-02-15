package lib

import (
	"encoding/binary"
	"time"
	"wasmgo/types"
	"wasmgo/wasm"
)

type Time struct {
	types.VMInterface
}

//long current_unix_time();
func (*Time) Current_unix_time() int64 {
	return time.Now().Unix()
}

//char * current_unix_time(long);
func (*Time) Unix_time_to(t int64) string {
	tm := time.Unix(t, 0)
	return tm.String()
}

//unsigned long clock()
//TODO
func (*Time) Clock() int {
	return time.Now().Nanosecond()
}

//time_t time(time_t *t)
func (ti *Time) Time(t int64) int64 {
	if t != 0 {
		binary.LittleEndian.PutUint64(ti.Vm.Memory[t:t+8], uint64(time.Now().Unix()))
	}
	return time.Now().Unix()
}

//char *ctime(const time_t *timer)
func (ti *Time) Ctime(t int64) int64 {
	tim := binary.LittleEndian.Uint64(ti.Vm.Memory[t : t+8])
	dataTimeStr := time.Unix(int64(tim), 0).Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	p := wasm.GetVMemory().Malloc(int64(len(dataTimeStr)))
	copy(ti.Vm.Memory[p:], []byte(dataTimeStr))
	return p
}

func (*Time) Gmtime(t int64) int64 {
	return 0
}
func (*Time) Asctime(t int64) int64 {
	return 0
}
