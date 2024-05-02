package encoding

import "fmt"

// 从切片 x(~[]T) 读取或向切片 x(~[]T) 写入数据，
// 并以 IO 操作流 operator 为媒介编码或解码数据。
// M 指代封装有在 T 上实现了基于 IO 操作流的 解码/编码 接口
func Slice[
	T any,
	S ~*[]T,
	M Marshal[T],
](operator IO, x S) error {
	len := len(*x)
	// prepare
	_, success := operator.(*Writer)
	if success {
		if len > MaxLenString {
			return fmt.Errorf("Slice: The length of the target slice is out of the max limited %d; len = %d", MaxLenSlice, len)
		}
		*x = make([]T, len)
	}
	// make slice if meet write mode
	length := uint32(len)
	operator.Uint32(&length)
	// write or read length
	for i := 0; i < len; i++ {
		err := M(&(*x)[i]).Marshal(operator)
		if err != nil {
			return fmt.Errorf("Slice: %v", err)
		}
	}
	// marshal or unmarshal data included
	return nil
	// return
}

// 从切片 x(~[]T) 读取或向切片 x(~[]T) 写入数据，
// 并以 IO 操作流 operator 为媒介编码或解码数据。
// marshal_func 指代基于 IO 操作流的 解码/编码 而实现的函数。

// 此函数与 Slice 的不同之处在于，此处要求直接提供对应的函数
// 用于切片的 解码/编码 ，因此此函数被认为广泛在基本数据类型的切片上使用
func SliceByFunc[T any, S ~*[]T](
	operator IO, x S,
	marshal_func func(x *T) error,
) error {
	len := len(*x)
	// prepare
	_, success := operator.(*Writer)
	if success {
		if len > MaxLenString {
			return fmt.Errorf("Slice: The length of the target slice is out of the max limited %d; len = %d", MaxLenSlice, len)
		}
		*x = make([]T, len)
	}
	// make slice if meet write mode
	length := uint32(len)
	operator.Uint32(&length)
	// write or read length
	for i := 0; i < len; i++ {
		err := marshal_func(&(*x)[i])
		if err != nil {
			return fmt.Errorf("SliceByFunc: %v", err)
		}
	}
	// marshal or unmarshal data included
	return nil
	// return
}

/*
从映射 x(map[K]V) 读取或向映射 x(map[K]V) 写入数据，
并以 IO 操作流 operator 为媒介编码或解码数据。

M 指代封装有在 V 上实现了基于 IO 操作流的 解码/编码 接口，
这被用于映射中值的 解码/编码 ；
key_marshal_func 指代基于 IO 操作流的 解码/编码 而实现的函数，
这被用于映射中键的 解码/编码 ；

key_marshal_func 与 M 的不同之处在于，前者要求直接提供对应的函数
用于切片的 解码/编码 ，因此映射的键在通常情况下应该是基本数据类型
*/
func Map[K comparable, V any, M Marshal[V]](
	operator IO,
	x map[K]V,
	key_marshal_func func(x *K) error,
) error {
	len := len(x)
	// prepare
	switch operator.(type) {
	case *Writer:
		if len > MapLenMap {
			return fmt.Errorf("Map: The length of the target map is out of the max limited %d; len = %d", MaxLenSlice, len)
		}
		// check length
		length := uint16(len)
		operator.Uint16(&length)
		// write length
		for key, value := range x {
			err := key_marshal_func(&key)
			if err != nil {
				return fmt.Errorf("Map: %v", err)
			}
			// write key
			err = M(&value).Marshal(operator)
			if err != nil {
				return fmt.Errorf("Map: %v", err)
			}
			// write value
		}
		// write data
	case *Reader:
		var length uint16
		operator.Uint16(&length)
		// read length
		for i := uint16(0); i < length; i++ {
			var key K
			err := key_marshal_func(&key)
			if err != nil {
				return fmt.Errorf("Map: %v", err)
			}
			// read key
			var value V
			err = M(&value).Marshal(operator)
			if err != nil {
				return fmt.Errorf("Map: %v", err)
			}
			// read value
			x[key] = value
			// sync data
		}
		// read data
	}
	// write or read data
	return nil
	// return
}
