package encoding

import (
	"encoding/binary"
	"fmt"
)

// 向写入者写入字符串 x
func (w *Writer) String(x *string) {
	if len(*x) > MaxLenString {
		panic(fmt.Sprintf("(w *Writer) String: The length of the target string is out of the max limited %d", MaxLenString))
	}
	// check length
	err := binary.Write(w.w, binary.BigEndian, uint16(len(*x)))
	if err != nil {
		panic(fmt.Sprintf("(w *Writer) String: %v", err))
	}
	// write the length of the target string
	err = w.WriteBytes([]byte(*x))
	if err != nil {
		panic(fmt.Sprintf("(w *Writer) String: %v", err))
	}
	// write string
}

// 向写入者写入布尔值 x
func (w *Writer) Bool(x *bool) {
	var err error
	// prepare
	switch *x {
	case false:
		err = w.WriteBytes([]byte{0})
	default:
		err = w.WriteBytes([]byte{1})
	}
	// write bool
	if err != nil {
		panic(fmt.Sprintf("(w *Writer) Bool: %v", err))
	}
	// return
}

// 向写入者写入 x(uint8)
func (w *Writer) Uint8(x *uint8) {
	if err := w.WriteBytes([]byte{*x}); err != nil {
		panic(fmt.Sprintf("(w *Writer) Uint8: %v", err))
	}
}

// 向写入者写入 x(int8)
func (w *Writer) Int8(x *int8) {
	if err := binary.Write(w.w, binary.BigEndian, *x); err != nil {
		panic(fmt.Sprintf("(w *Writer) Int8: %v", err))
	}
}

// 向写入者写入 x(uint16)
func (w *Writer) Uint16(x *uint16) {
	if err := binary.Write(w.w, binary.BigEndian, *x); err != nil {
		panic(fmt.Sprintf("(w *Writer) Uint16: %v", err))
	}
}

// 向写入者写入 x(int16)
func (w *Writer) Int16(x *int16) {
	if err := binary.Write(w.w, binary.BigEndian, *x); err != nil {
		panic(fmt.Sprintf("(w *Writer) Int16: %v", err))
	}
}

// 向写入者写入 x(uint32)
func (w *Writer) Uint32(x *uint32) {
	if err := binary.Write(w.w, binary.BigEndian, *x); err != nil {
		panic(fmt.Sprintf("(w *Writer) Uint32: %v", err))
	}
}

// 向写入者写入 x(int32)
func (w *Writer) Int32(x *int32) {
	if err := binary.Write(w.w, binary.BigEndian, *x); err != nil {
		panic(fmt.Sprintf("(w *Writer) Int32: %v", err))
	}
}

// 向写入者写入 x(uint64)
func (w *Writer) Uint64(x *uint64) {
	if err := binary.Write(w.w, binary.BigEndian, *x); err != nil {
		panic(fmt.Sprintf("(w *Writer) Uint64: %v", err))
	}
}

// 向写入者写入 x(int64)
func (w *Writer) Int64(x *int64) {
	if err := binary.Write(w.w, binary.BigEndian, *x); err != nil {
		panic(fmt.Sprintf("(w *Writer) Int64: %v", err))
	}
}
