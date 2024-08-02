package math

import "unsafe"

const INT8 int8 = 1
const INT16 int16 = 1
const INT32 int32 = 1
const INT64 int64 = 1

func SizeofInt8() uintptr {
	return unsafe.Sizeof(INT8)
}

func SizeofInt16() uintptr {
	return unsafe.Sizeof(INT16)
}

func SizeofInt32() uintptr {
	return unsafe.Sizeof(INT32)
}

func SizeofInt64() uintptr {
	return unsafe.Sizeof(INT64)
}
