package lib

import (
	"wasmgo/types"
)

type MemoryInterface struct {
	types.VMInterface
}

///////////////////////////////////////////////////////////////////////////////

//将s中当前位置后面的n个字节（typedef unsigned int size_t ）用 ch 替换并返回 s 。
func (m *MemoryInterface) Memset(s int64, ch int, n int) int64 {
	for i := 0; i < n; i++ {
		m.Vm.Memory[s+int64(i)] = byte(ch)
	}
	return s
}

//void *memcpy(void *str1, const void *str2, size_t n) 从存储区 src 复制 n 个字符到存储区 des。
func (m *MemoryInterface) Memcpy(des int64, src int64, n int) int64 {
	copy(m.Vm.Memory[des:des+int64(n)], m.Vm.Memory[src:src+int64(n)])
	return des

}
