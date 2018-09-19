package main

// GetFieldInt do smth...
func (this *Int) GetFieldInt() int {
	if this != nil {
		return this.FieldInt
	}
	return 0
}

// GetFieldInt8 do smth...
func (this *Int) GetFieldInt8() int8 {
	if this != nil {
		return this.FieldInt8
	}
	return 0
}

// GetFieldInt16 do smth...
func (this *Int) GetFieldInt16() int16 {
	if this != nil {
		return this.FieldInt16
	}
	return 0
}

// GetFieldInt32 do smth...
func (this *Int) GetFieldInt32() int32 {
	if this != nil {
		return this.FieldInt32
	}
	return 0
}

// GetFieldInt64 do smth...
func (this *Int) GetFieldInt64() int64 {
	if this != nil {
		return this.FieldInt64
	}
	return 0
}

// GetFieldUint do smth...
func (this *Uint) GetFieldUint() uint {
	if this != nil {
		return this.FieldUint
	}
	return 0
}

// GetFieldUint8 do smth...
func (this *Uint) GetFieldUint8() uint8 {
	if this != nil {
		return this.FieldUint8
	}
	return 0
}

// GetFieldUint16 do smth...
func (this *Uint) GetFieldUint16() uint16 {
	if this != nil {
		return this.FieldUint16
	}
	return 0
}

// GetFieldUint32 do smth...
func (this *Uint) GetFieldUint32() uint32 {
	if this != nil {
		return this.FieldUint32
	}
	return 0
}

// GetFieldUint64 do smth...
func (this *Uint) GetFieldUint64() uint64 {
	if this != nil {
		return this.FieldUint64
	}
	return 0
}

// GetFieldUintPtr do smth...
func (this *Uint) GetFieldUintPtr() uintptr {
	if this != nil {
		return this.FieldUintPtr
	}
	return 0
}

// GetFieldFloat32 do smth...
func (this *Float) GetFieldFloat32() float32 {
	if this != nil {
		return this.FieldFloat32
	}
	return 0.0
}

// GetFieldFloat64 do smth...
func (this *Float) GetFieldFloat64() float64 {
	if this != nil {
		return this.FieldFloat64
	}
	return 0.0
}

// GetFieldComplex64 do smth...
func (this *Complex) GetFieldComplex64() complex64 {
	if this != nil {
		return this.FieldComplex64
	}
	return (0 + 0i)
}

// GetFieldComplex128 do smth...
func (this *Complex) GetFieldComplex128() complex128 {
	if this != nil {
		return this.FieldComplex128
	}
	return (0 + 0i)
}

// GetFieldString do smth...
func (this *Other) GetFieldString() string {
	if this != nil {
		return this.FieldString
	}
	return ""
}

// GetFieldBool do smth...
func (this *Other) GetFieldBool() bool {
	if this != nil {
		return this.FieldBool
	}
	return false
}

// GetFieldByte do smth...
func (this *Other) GetFieldByte() byte {
	if this != nil {
		return this.FieldByte
	}
	return 0
}

// GetFieldRune do smth...
func (this *Other) GetFieldRune() rune {
	if this != nil {
		return this.FieldRune
	}
	return 0
}

// GetFieldSub do smth...
func (this *SubStruct) GetFieldSub() Sub {
	if this != nil {
		return this.FieldSub
	}
	return Sub{}
}

// GetFieldSubPtr do smth...
func (this *SubStruct) GetFieldSubPtr() *Sub {
	if this != nil {
		return this.FieldSubPtr
	}
	return &Sub{}
}

// GetFieldSlice do smth...
func (this *Slice) GetFieldSlice() []int {
	if this != nil {
		return this.FieldSlice
	}
	return []int{}
}

// GetFieldSliceSub do smth...
func (this *Slice) GetFieldSliceSub() []Sub {
	if this != nil {
		return this.FieldSliceSub
	}
	return []Sub{}
}

// GetFieldSliceSubPtr do smth...
func (this *Slice) GetFieldSliceSubPtr() []*Sub {
	if this != nil {
		return this.FieldSliceSubPtr
	}
	return []*Sub{}
}

// GetFieldMapIntInt do smth...
func (this *Map) GetFieldMapIntInt() map[int]int {
	if this != nil {
		return this.FieldMapIntInt
	}
	return map[int]int{}
}

// GetFieldMapIntSub do smth...
func (this *Map) GetFieldMapIntSub() map[int]Sub {
	if this != nil {
		return this.FieldMapIntSub
	}
	return map[int]Sub{}
}

// GetFieldMapIntSubPtr do smth...
func (this *Map) GetFieldMapIntSubPtr() map[int]*Sub {
	if this != nil {
		return this.FieldMapIntSubPtr
	}
	return map[int]*Sub{}
}
