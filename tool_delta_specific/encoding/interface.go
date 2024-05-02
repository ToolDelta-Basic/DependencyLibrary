package encoding

import (
	"bytes"
)

// 为二进制数据实现的 IO 操作流
type IO interface {
	AutoMarshal
	GetBuffer
}

// 取得阅读器或写入者的底层切片
type GetBuffer interface {
	GetBuffer() *bytes.Buffer
}

// 指代在 *T 上实现了基于 IO 操作流的 解码/编码 接口
type Marshal[T any] interface {
	Marshal(operator IO) error
	*T
}

// 为二进制数据实现的自动化 解码/编码 实现。
//
// 以下列出的每个函数都提供了两个实现，
// 以允许编码或解码二进制数据。
//
// 当传入 encoding.Reader 时，数据将从 encoding.Reader 解码至 x ；
// 当传入 encoding.Writer 时，x 将被编码至 encoding.Writer
type AutoMarshal interface {
	String(x *string) error
	Bool(x *bool) error
	Uint8(x *uint8) error
	Int8(x *int8) error
	Uint16(x *uint16) error
	Int16(x *int16) error
	Uint32(x *uint32) error
	Int32(x *int32) error
	Uint64(x *uint64) error
	Int64(x *int64) error
}
