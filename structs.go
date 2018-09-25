package main

import (
	"time"
)

// all the structs below are supported
// s with type interface are not supported

type Int struct {
	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64
}

type Uint struct {
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	UintPtr uintptr
}

type Float struct {
	Float32 float32
	Float64 float64
}

type Complex struct {
	Complex64  complex64
	Complex128 complex128
}

type Other struct {
	String string
	Bool   bool
	// alias for uint8
	Byte byte
	// alias for int32
	Rune rune
}

type Struct struct {
	Sub    Sub
	SubPtr *Sub
}

type Sub struct{}

type Slice struct {
	Slice       []int
	SliceSub    []Sub
	SliceSubPtr []*Sub
}

type Array struct {
	Array       [10]int
	ArraySub    [10]Sub
	ArraySubPtr [10]*Sub
}

type Map struct {
	MapIntInt    map[int]int
	MapIntSub    map[int]Sub
	MapIntSubPtr map[int]*Sub
	IntSliceSub  [][]*Sub
}

type Time struct {
	Time    time.Time
	TimePtr *time.Time
}

type Custom struct {
	// shouldn't be a getter for this field
	Custom  interface{}
	Custom2 map[int][][][]map[float32][]*Sub
	Custom3 [][][2][][10]map[int]*time.Time
}
