package lib

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"wasmgo/types"
)

type StdLib struct {
	types.RegInterface
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

//int sbrk(void *addr);
//func (*StdLib) Sbrk(increment int) int {
//	log.Printf("############# %d",increment)
//	//increment = increment | 0
//	//var oldDynamicTop = 0;
//	//var oldDynamicTopOnChange = 0;
//	//var newDynamicTop = 0;
//	//var totalMemory = 0;
//	return -1
//}

//char *getenv(char *envvar);
func (*StdLib) Getenv(envvar string) string {
	return os.Getenv(envvar)
}
func (*StdLib) BuildEnvironment(environ int64) {
	//fmt.Println(environ)
}
