package djb2

import (
	"unsafe"
)
// bernstein's djb2 hash
func Hash(s string) uint32 {
	var h uint32 = 5381
	for i := 0; i < len(s); i++ {
		h = 33*h + uint32(s[i])
	}
	return h
}

// unsafe.Slice((*uint32)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)), len(s))
func Hash2(s string) uint32 {
	var h uint32 = 5381
	ss := *(*[]uint32)(unsafe.Pointer(&s))
	for i := 0; i < len(ss); i++ {
		h = (h*33) ^ ss[i]
	}
	return h
}
