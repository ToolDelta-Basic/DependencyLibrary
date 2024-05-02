package encoding

import (
	"io"
)

// 用于读取二进制切片的阅读器
type Reader struct {
	r interface{ io.Reader }
}

// 用于写入二进制切片的写入者
type Writer struct {
	w interface{ io.Writer }
}
