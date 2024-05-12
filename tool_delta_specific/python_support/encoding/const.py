# 描述一些基本数据类型中，
# 其中可含有元素的最大数目
MAX_LEN_STRING: int = 65535  # 单个字符串的最大长度上限 (uint16)
MAX_LEN_SLICE: int = 4294967295  # 单个切片的最大长度上限 (uint32)
MAX_LEN_MAP: int = 65535  # 单个映射的最大长度上限 (uint16)

# 描述长度固定的数据的长度
CONST_LEN_UUID: int = 16  # 单个 uuid.UUID 的固定长度
