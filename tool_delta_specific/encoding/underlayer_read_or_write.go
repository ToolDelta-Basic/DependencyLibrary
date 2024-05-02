package encoding

import "fmt"

// 向写入者写入字节切片 p
func (w *Writer) WriteBytes(p []byte) error {
	_, err := w.w.Write(p)
	if err != nil {
		return fmt.Errorf("WriteBytes: %v", err)
	}
	return nil
}

// 从阅读器阅读 length 个字节
func (r *Reader) ReadBytes(length int) ([]byte, error) {
	ans := make([]byte, length)
	_, err := r.r.Read(ans)
	if err != nil {
		return nil, fmt.Errorf("ReadBytes: %v", err)
	}
	return ans, nil
}
