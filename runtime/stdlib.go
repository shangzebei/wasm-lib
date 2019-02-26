package lib

import (
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"wasmgo/types"
	"wasmgo/wasm"
)

type StdLib struct {
	types.RegInterface
	types.VMInterface
}

func (s *StdLib) Init() {
	s.Replace("BuildEnvironment", "__buildEnvironment")

}

//double atof(const char *str)
func (*StdLib) Atof(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

//int atoi(const char *str)
func (*StdLib) Atoi(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		println(e)
	}
	return i
}

//long int atol(const char *str)
func (*StdLib) Atol(str string) int64 {
	f, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

//void exit(int status)
func (*StdLib) Exit(status int) {
	os.Exit(status)
}

//int abs(int x)
func (*StdLib) Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

//double fabs(double x)
func (*StdLib) Fabs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (*StdLib) Sleep(t int64) {
	time.Sleep(time.Duration(time.Nanosecond.Nanoseconds() * t))
}

//void abort(void)
func (*StdLib) Abort() {
	os.Exit(0)
}

//char *getenv(char *envvar);
func (*StdLib) Getenv(envvar string) string {
	return os.Getenv(envvar)
}

func (*StdLib) BuildEnvironment(environ int64) {
	//fmt.Println(environ)
}

//void (*signal (int sig, void (*func)(int)))(int);
func (s *StdLib) Signal(sig int, functionId int) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.Signal(sig), syscall.SIGTERM)
	go func() {
		<-c
		wasm.InvokeMethod(s.Vm, functionId, int64(sig))
	}()
}

/*struct timespec
{
  time_t  tv_sec;          秒seconds
  long    tv_nsec;         纳秒nanoseconds
};
*/

//int nanosleep(const struct timespec *req, struct timespec *rem);
func (s *StdLib) Nanosleep(req int64, rem int64) int32 {
	sec := binary.LittleEndian.Uint32(s.Vm.Memory[req : req+4])
	req += 4
	nanosec := binary.LittleEndian.Uint32(s.Vm.Memory[req : req+4])
	time.Sleep(time.Duration(time.Second.Nanoseconds()*int64(sec) + int64(nanosec)))
	return 0
}

//int raise (signal sig);
func (*StdLib) Raise(sig int64) int32 {
	fmt.Println("Raise not implement")
	return 0
}
