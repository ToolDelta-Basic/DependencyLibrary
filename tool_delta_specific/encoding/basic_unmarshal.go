package encoding

import (
	"encoding/binary"
	"fmt"
)

// 从阅读器阅读一个字符串并返回到 x 上
func (r *Reader) String(x *string) error {
	var length uint16
	err := binary.Read(r.r, binary.BigEndian, &length)
	if err != nil {
		return fmt.Errorf("(r *Reader) String: %v", err)
	}
	// get the length of the target string
	stringBytes, err := r.ReadBytes(int(length))
	if err != nil {
		return fmt.Errorf("(r *Reader) String: %v", err)
	}
	*x = string(stringBytes)
	// get the target string
	return nil
	// return
}

// 从阅读器阅读一个布尔值并返回到 x 上
func (r *Reader) Bool(x *bool) error {
	ans, err := r.ReadBytes(1)
	if err != nil {
		return fmt.Errorf("(r *Reader) Bool: %v", err)
	}
	switch ans[0] {
	case 0:
		*x = false
	case 1:
		*x = true
	case 2:
		return fmt.Errorf("(r *Reader) Bool: Unexpected value %#v was find", ans)
	}
	return nil
}

// 从阅读器阅读一个 uint8 并返回到 x 上
func (r *Reader) Uint8(x *uint8) error {
	ans, err := r.ReadBytes(1)
	if err != nil {
		return fmt.Errorf("(r *Reader) Uint8: %v", err)
	}
	*x = ans[0]
	return nil
}

// 从阅读器阅读一个 int8 并返回到 x 上
func (r *Reader) Int8(x *int8) error {
	err := binary.Read(r.r, binary.BigEndian, x)
	if err != nil {
		return fmt.Errorf("(r *Reader) Int8: %v", err)
	}
	return nil
}

// 从阅读器阅读一个 uint16 并返回到 x 上
func (r *Reader) Uint16(x *uint16) error {
	err := binary.Read(r.r, binary.BigEndian, x)
	if err != nil {
		return fmt.Errorf("(r *Reader) Uint16: %v", err)
	}
	return nil
}

// 从阅读器阅读一个 int16 并返回到 x 上
func (r *Reader) Int16(x *int16) error {
	err := binary.Read(r.r, binary.BigEndian, x)
	if err != nil {
		return fmt.Errorf("(r *Reader) Int16: %v", err)
	}
	return nil
}

// 从阅读器阅读一个 uint32 并返回到 x 上
func (r *Reader) Uint32(x *uint32) error {
	err := binary.Read(r.r, binary.BigEndian, x)
	if err != nil {
		return fmt.Errorf("(r *Reader) Uint32: %v", err)
	}
	return nil
}

// 从阅读器阅读一个 int32 并返回到 x 上
func (r *Reader) Int32(x *int32) error {
	err := binary.Read(r.r, binary.BigEndian, x)
	if err != nil {
		return fmt.Errorf("(r *Reader) Int32: %v", err)
	}
	return nil
}

// 从阅读器阅读一个 uint64 并返回到 x 上
func (r *Reader) Uint64(x *uint64) error {
	err := binary.Read(r.r, binary.BigEndian, x)
	if err != nil {
		return fmt.Errorf("(r *Reader) Uint64: %v", err)
	}
	return nil
}

// 从阅读器阅读一个 int64 并返回到 x 上
func (r *Reader) Int64(x *int64) error {
	err := binary.Read(r.r, binary.BigEndian, x)
	if err != nil {
		return fmt.Errorf("(r *Reader) Int64: %v", err)
	}
	return nil
}
