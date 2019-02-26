package lib

import (
	"encoding/binary"
	"log"
	"math"
	"os"
	"wasmgo/types"
	"wasmgo/wasm"
)

type SystemCall struct {
	types.VMInterface
	types.RegInterface
	varargs int32
	Streams map[int]*os.File
	Buffers map[int]interface{}
}

func (s *SystemCall) Init() {
	s.Buffers = make(map[int]interface{})
	s.Buffers[0] = nil
	s.Buffers[1] = make([]byte, 0)
	s.Buffers[2] = make([]byte, 0)

	s.Replace("Syscall6", "__syscall6")   // close
	s.Replace("Syscall54", "__syscall54") //ioctl
	s.Replace("Syscall140", "__syscall140")
	s.Replace("Syscall146", "__syscall146")
	s.Replace("Cxa_atexit", "__cxa_atexit")
	s.Replace("Syscall5", "__syscall5")
	s.Replace("Syscall145", "__syscall145")
	s.Replace("Syscall38", "__syscall38")
	s.Replace("Syscall10", "__syscall10")

	//var FuncList = make(map[string]interface{})
	s.Streams = make(map[int]*os.File)
	s.Streams[1] = os.Stdout
	s.Streams[2] = os.Stderr

}

//close
func (s *SystemCall) Syscall6(a int, b int) int {
	// close
	err := s.Streams[a].Close()
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Syscall6 a = %d,b =%d \n", a, b)
	return 0
}

func (*SystemCall) Cxa_atexit() {
	log.Printf("atexit() called, but EXIT_RUNTIME is not set, so atexits() will not be called. set EXIT_RUNTIME to 1 (see the FAQ)")
}

//ioctl
func (s *SystemCall) Syscall54(a int32, b int32) int32 {
	s.varargs = b
	//fmt.Printf("Syscall54 a = %d,b =%d \n", a, b)
	return 0
}

//llseek
func (s *SystemCall) Syscall140(a int32, b int32) int32 {
	s.varargs = b
	stream := s.get()
	s.get()               // NOTE: offset_high is unused - Emscripten's off_t is 32-bit
	offset_low := s.get() //var offset = offset_low;
	result := s.get()
	whence := s.get()
	//fmt.Println(stream, offset_high, offset_low, result, whence)
	return s.llseek(stream, int64(offset_low), whence, result)
}

//read
func (s *SystemCall) Syscall145(a int32, b int32) int {
	s.varargs = b
	stream := s.get()
	iov := s.get()
	iovcnt := s.get()
	//fmt.Println(stream, iov, iovcnt)
	var ret = 0
	pos := iov
	for i := 0; i < iovcnt; i++ {
		ptr := int(binary.LittleEndian.Uint32(s.Vm.Memory[pos : pos+4]))
		len := int(binary.LittleEndian.Uint32(s.Vm.Memory[pos+4 : pos+8]))
		pos += 8
		ret += s.read(stream, ptr, len)
	}
	return ret
}

// open
func (s *SystemCall) Syscall5(a int32, b int32) int64 {
	s.varargs = b
	fp := s.get()
	filename := wasm.GetString(int64(fp), s.Vm)
	flags := s.get()
	mode := s.get()
	file, _ := os.OpenFile(filename, flags, os.FileMode(mode))
	s.Streams[int(file.Fd())] = file
	s.Buffers[int(file.Fd())] = make([]byte, 0)
	//fmt.Println(error)
	return int64(file.Fd())
}

//write
func (s *SystemCall) Syscall146(a int32, b int32) int {
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
			s.write(stream, int(s.Vm.Memory[ptr+j]))
		}
		pos += 8
		ret = ret + len
	}
	return ret
}

//rename
func (s *SystemCall) Syscall38(a int32, b int32) int32 {
	s.varargs = b
	old_path := wasm.GetString(int64(s.get()), s.Vm)
	new_path := wasm.GetString(int64(s.get()), s.Vm)
	err := os.Rename(old_path, new_path)
	if err != nil {
		log.Println(err)
		return -1
	}
	return 0
}

//remove
func (s *SystemCall) Syscall10(a int32, b int32) int32 {
	s.varargs = b
	path := wasm.GetString(int64(s.get()), s.Vm)
	err := os.Remove(path)
	if err != nil {
		log.Println(err)
		return -1
	}
	return 0
}

func (s *SystemCall) get() int {
	ret := s.Vm.Memory[s.varargs : s.varargs+4]
	s.varargs += 4
	return int(binary.LittleEndian.Uint32(ret))

}

func (s *SystemCall) write(stream int, curr int) {
	var buff = s.Buffers[stream].([]byte)
	buff = append(buff, byte(curr))
	if curr == 0 || curr == 10 {
		_, error := s.Streams[stream].Write(buff)
		if error != nil {
			log.Println(error)
		}
		buff = nil
	}
	s.Buffers[stream] = buff

}

func (s *SystemCall) read(stream int, ptr int, lenght int) int {
	b := make([]byte, lenght)
	st := s.Streams[stream]
	n, _ := st.Read(b)
	copy(s.Vm.Memory[ptr:ptr+lenght], b)
	return n
}

func (s *SystemCall) llseek(stream int, offset int64, whence int, resultPtr int) int32 {
	ret, error := s.Streams[stream].Seek(offset, whence)
	if error != nil {
		log.Println(error)
	}
	if ret > math.MaxInt32 {
		log.Printf("over %d", ret)
	}
	binary.LittleEndian.PutUint32(s.Vm.Memory[resultPtr:resultPtr+4], uint32(ret))
	return int32(ret)
}
