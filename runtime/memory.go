package lib

import (
	"wasmgo/types"
	"wasmgo/wasm"
)

type MemoryInterface struct {
	types.VMInterface
}

///////////////////////////////////////////////////////////////////////////////

//将s中当前位置后面的n个字节（typedef unsigned int size_t ）用 ch 替换并返回 s 。
func (m *MemoryInterface) Memset(s int64, ch int, n int) int64 {
	if s < 0 {
		s = wasm.FindfreeSpece(m.Vm, int64(n))
	}
	for i := 0; i < n; i++ {
		m.Vm.Memory[s+int64(i)] = byte(ch)
	}
	return s
}

//void *memcpy(void *str1, const void *str2, size_t n) 从存储区 src 复制 n 个字符到存储区 des。
func (m *MemoryInterface) Memcpy(des int64, src int64, n int) int64 {
	if des > 0 {
		end := wasm.GetMemoryEndPos(des, m.Vm)
		if (end - des) < int64(n) {
			panic("des Out of memory")
		}
		copy(m.Vm.Memory[des:des+int64(n)], m.Vm.Memory[src:src+int64(n)])
		return des
	} else {
		p := m.Malloc(int64(n))
		copy(m.Vm.Memory[p:p+int64(n)], m.Vm.Memory[src:src+int64(n)])
		return p
	}

}

//void *pTemp = malloc(nSize);
func (m *MemoryInterface) Malloc(size int64) int64 {
	return wasm.FindfreeSpece(m.Vm, size)
}

//void free(void *);
func (m *MemoryInterface) Free(p int64) {
	end := wasm.GetMemoryEndPos(p, m.Vm)
	m.Memset(p, 0, int(end-p))
}
