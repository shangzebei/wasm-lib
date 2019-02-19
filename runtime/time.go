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
	defer func() {
		wasm.GetVMemory().Free(p)
	}()
	return p
}

/*
struct tm {
	int	tm_sec;		/* seconds after the minute [0-60]
	int	tm_min;		/* minutes after the hour [0-59]
	int	tm_hour;	/* hours since midnight [0-23]
	int	tm_mday;	/* day of the month [1-31]
	int	tm_mon;		/* months since January [0-11]
	int	tm_year;	/* years since 1900
	int	tm_wday;	/* days since Sunday [0-6]
	int	tm_yday;	/* days since January 1 [0-365]
	int	tm_isdst;	/* Daylight Savings Time flag
	long	tm_gmtoff;	/* offset from UTC in seconds
	char	*tm_zone;	/* timezone abbreviation
};
*/
//struct tm *gmtime(const time_t *timer)
func (ti *Time) Gmtime(t int64) int64 {
	tim := binary.LittleEndian.Uint64(ti.Vm.Memory[t : t+8])
	op := wasm.GetVMemory().Malloc(44)
	tAll := time.Unix(int64(tim), 0)
	p := op
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.Second())) //tm_sec
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.Minute())) //tm_min
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.Hour())) //tm_hour
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.Day())) //tm_mday
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.Month())) //tm_mon
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.Year())) //tm_year
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.Weekday())) //tm_wday
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(tAll.YearDay())) //tm_yday
	p += 4
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(0)) //tm_isdst
	p += 4
	sz, off := tAll.Zone()
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(off)) //tm_gmtoff
	p += 4
	sp := wasm.GetVMemory().Malloc(int64(len(sz)))
	copy(ti.Vm.Memory[sp:sp+int64(len(sz))], []byte(sz))
	binary.LittleEndian.PutUint32(ti.Vm.Memory[p:p+4], uint32(sp)) //*tm_zone
	return op
}
func (*Time) Asctime(t int64) int64 {
	return 0
}
