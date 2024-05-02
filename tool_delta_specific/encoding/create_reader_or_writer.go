package encoding

import (
	"bytes"
)

// 创建一个新的阅读器
func NewReader(reader *bytes.Buffer) IO {
	return &Reader{r: reader}
}

// 创建一个新的写入者
func NewWriter(writer *bytes.Buffer) IO {
	return &Writer{w: writer}
}

// 取得阅读器的底层切片
func (r *Reader) GetBuffer() *bytes.Buffer {
	return r.r.(*bytes.Buffer)
}

// 取得写入者的底层切片
func (w *Writer) GetBuffer() *bytes.Buffer {
	return w.w.(*bytes.Buffer)
}
