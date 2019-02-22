package lib

import (
	"fmt"
	"os/exec"
)

type System struct {
}

//void * dlopen(const char * __path, int __mode);
func (s *System) Dlopen(path string, mode int32) int64 {

	return 1
}

//void * dlsym(void * __handle, const char * __symbol);
func (s *System) Dlsym(handle int64, symbol string) int64 {
	return 1
}

//char * dlerror(void);
func (s *System) Dlerror() int64 {
	return 0
}

//int dlclose(void * __handle);
func (s *System) Dlclose(handle int64) int64 {
	return 0
}

//int system(const char *)
func (s *System) System(string2 string) int {
	cmd := exec.Command(string2)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	return 0
}
