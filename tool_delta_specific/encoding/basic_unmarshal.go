package encoding

import (
	"encoding/binary"
	"fmt"
)

// 从阅读器阅读一个字符串并返回到 x 上
func (r *Reader) String(x *string) {
	var length uint16
	if err := binary.Read(r.r, binary.BigEndian, &length); err != nil {
		panic(fmt.Sprintf("(r *Reader) String: %v", err))
	}
	// get the length of the target string
	string_bytes, err := r.ReadBytes(int(length))
	if err != nil {
		panic(fmt.Sprintf("(r *Reader) String: %v", err))
	}
	*x = string(string_bytes)
	// get the target string
}

// 从阅读器阅读一个布尔值并返回到 x 上
func (r *Reader) Bool(x *bool) {
	ans, err := r.ReadBytes(1)
	if err != nil {
		panic(fmt.Sprintf("(r *Reader) Bool: %v", err))
	}
	// read one byte
	switch ans[0] {
	case 0:
		*x = false
	case 1:
		*x = true
	default:
		panic(fmt.Sprintf("(r *Reader) Bool: Unexpected value %d was find", ans[0]))
	}
	// set data
}

// 从阅读器阅读一个 uint8 并返回到 x 上
func (r *Reader) Uint8(x *uint8) {
	ans, err := r.ReadBytes(1)
	if err != nil {
		panic(fmt.Sprintf("(r *Reader) Uint8: %v", err))
	}
	*x = ans[0]
}

// 从阅读器阅读一个 int8 并返回到 x 上
func (r *Reader) Int8(x *int8) {
	if err := binary.Read(r.r, binary.BigEndian, x); err != nil {
		panic(fmt.Sprintf("(r *Reader) Int8: %v", err))
	}
}

// 从阅读器阅读一个 uint16 并返回到 x 上
func (r *Reader) Uint16(x *uint16) {
	if err := binary.Read(r.r, binary.BigEndian, x); err != nil {
		panic(fmt.Sprintf("(r *Reader) Uint16: %v", err))
	}
}

// 从阅读器阅读一个 int16 并返回到 x 上
func (r *Reader) Int16(x *int16) {
	if err := binary.Read(r.r, binary.BigEndian, x); err != nil {
		panic(fmt.Sprintf("(r *Reader) Int16: %v", err))
	}
}

// 从阅读器阅读一个 uint32 并返回到 x 上
func (r *Reader) Uint32(x *uint32) {
	if err := binary.Read(r.r, binary.BigEndian, x); err != nil {
		panic(fmt.Sprintf("(r *Reader) Uint32: %v", err))
	}
}

// 从阅读器阅读一个 int32 并返回到 x 上
func (r *Reader) Int32(x *int32) {
	if err := binary.Read(r.r, binary.BigEndian, x); err != nil {
		panic(fmt.Sprintf("(r *Reader) Int32: %v", err))
	}
}

// 从阅读器阅读一个 uint64 并返回到 x 上
func (r *Reader) Uint64(x *uint64) {
	if err := binary.Read(r.r, binary.BigEndian, x); err != nil {
		panic(fmt.Sprintf("(r *Reader) Uint64: %v", err))
	}
}

// 从阅读器阅读一个 int64 并返回到 x 上
func (r *Reader) Int64(x *int64) {
	if err := binary.Read(r.r, binary.BigEndian, x); err != nil {
		panic(fmt.Sprintf("(r *Reader) Int64: %v", err))
	}
}
