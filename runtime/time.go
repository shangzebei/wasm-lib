package lib

import "time"

type Time struct {
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
