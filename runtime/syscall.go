package lib

import (
	"encoding/binary"
	"fmt"
	"wasmgo/types"
)

type SystemCall struct {
	types.VMInterface
	types.RegInterface
	varargs int
	Buffers []interface{}
}

func (s *SystemCall) Init() {
	s.Buffers = make([]interface{}, 3)
	s.Buffers[0] = nil
	s.Buffers[1] = make([]byte, 0)
	s.Buffers[2] = make([]byte, 0)

	s.Replace("Syscall6", "__syscall6")   // close
	s.Replace("Syscall54", "__syscall54") //ioctl
	s.Replace("Syscall140", "__syscall140")
	s.Replace("Syscall146", "__syscall146")
	s.Replace("Cxa_atexit", "__cxa_atexit")
}

//extern int __syscall6(int a,int b);
func (*SystemCall) Syscall6(a int, b int) int {
	// close
	fmt.Printf("Syscall6 a = %d,b =%d \n", a, b)
	return 0
}

func (*SystemCall) Cxa_atexit() {
	fmt.Println("atexit() called, but EXIT_RUNTIME is not set, so atexits() will not be called. set EXIT_RUNTIME to 1 (see the FAQ)")
}

func (s *SystemCall) Syscall54(a int, b int) int {
	s.varargs = b
	// ioctl
	//fmt.Printf("Syscall54 a = %d,b =%d \n", a, b)
	return 0
}
func (s *SystemCall) Syscall140(a int, b int) int {
	fmt.Printf("Syscall140 a = %d,b =%d \n", a, b)
	s.varargs = b
	return 0
}
func (s *SystemCall) Syscall146(a int, b int) int {
	//fmt.Printf("Syscall146 a = %d  b =%d \n", a, b)
	s.varargs = b
	stream := s.get()
	iov := s.get()
	iovcnt := s.get()
	//fmt.Printf(" stream = %d  iov =%d iovcnt =%d \n", stream, iov, iovcnt)
	var ret = 0
	pos := iov
	for i := 0; i < iovcnt; i++ {
		ptr := int(binary.LittleEndian.Uint32(s.Vm.Memory[pos : pos+4]))
		len := int(binary.LittleEndian.Uint32(s.Vm.Memory[pos+4 : pos+8]))
		//fmt.Printf(" ptr = %d  len =%d  \n", ptr, len)
		for j := 0; j < len; j++ {
			s.printChar(stream, int(s.Vm.Memory[ptr+j]))
		}
		pos += 8
		ret = ret + len
	}
	return ret
}

func (s *SystemCall) get() int {
	ret := s.Vm.Memory[s.varargs : s.varargs+4]
	s.varargs += 4
	return int(binary.LittleEndian.Uint32(ret))

}

func (s *SystemCall) printChar(stream int, curr int) {
	var buff = s.Buffers[stream].([]byte)
	if curr == 0 || curr == 10 {
		fmt.Println(string(buff))
		buff = nil
	} else {
		buff = append(buff, byte(curr))
	}
	s.Buffers[stream] = buff
}
