package encoding

// 描述一些基本数据类型中，
// 其中可含有元素的最大数目
const (
	MaxLenString = 65535      // 单个字符串的最大长度上限
	MaxLenSlice  = 4294967295 // 单个切片的最大长度上限
	MapLenMap    = 65535      // 单个映射的最大长度上限
)

// 描述长度固定的数据的长度
const (
	ConstLenUUID = 16 // 单个 uuid.UUID 的固定长度
)
