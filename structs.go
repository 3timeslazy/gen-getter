package main

// structs for tests
// all the structs below are supported

type Int struct {
	FieldInt   int
	FieldInt8  int8
	FieldInt16 int16
	FieldInt32 int32
	FieldInt64 int64
}

type Uint struct {
	FieldUint    uint
	FieldUint8   uint8
	FieldUint16  uint16
	FieldUint32  uint32
	FieldUint64  uint64
	FieldUintPtr uintptr
}

type Float struct {
	FieldFloat32 float32
	FieldFloat64 float64
}

type Complex struct {
	FieldComplex64  complex64
	FieldComplex128 complex128
}

type Other struct {
	FieldString string
	FieldBool   bool
	// alias for uint8
	FieldByte byte
	// alias for int32
	FieldRune rune
}

type SubStruct struct {
	FieldSub    Sub
	FieldSubPtr *Sub
}

type Sub struct{}

type Slice struct {
	FieldSlice       []int
	FieldSliceSub    []Sub
	FieldSliceSubPtr []*Sub
}

type Map struct {
	FieldMapIntInt    map[int]int
	FieldMapIntSub    map[int]Sub
	FieldMapIntSubPtr map[int]*Sub
}
