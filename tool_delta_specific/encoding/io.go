package encoding

import "fmt"

// 从切片 x(~[]T) 读取或向切片 x(~[]T) 写入数据，
// 并以 IO 操作流 operator 为媒介编码或解码数据。
// M 指代封装有在 T 上实现了基于 IO 操作流的 解码/编码 接口
func Slice[
	T any,
	S ~*[]T,
	M Marshal[T],
](operator IO, x S) {
	len := len(*x)
	if len > MaxLenSlice {
		panic(fmt.Sprintf("Slice: The length of the target slice is out of the max limited %d; len = %d", MaxLenSlice, len))
	}
	// pre check
	length := uint32(len)
	operator.Uint32(&length)
	// wrire or read length
	_, success := operator.(*Reader)
	if success {
		*x = make([]T, length)
	}
	// make slice if meet read mode
	for i := uint32(0); i < length; i++ {
		M(&(*x)[i]).Marshal(operator)
	}
	// marshal or unmarshal data included
}

// 从切片 x(~[]T) 读取或向切片 x(~[]T) 写入数据，
// 并以 IO 操作流 operator 为媒介编码或解码数据。
// marshal_func 指代基于 IO 操作流而实现的函数，
// 其被用于切片中数据的 解码/编码 。

// 此函数与 Slice 的不同之处在于，
// 此处要求直接提供对应的函数用于切片的 解码/编码 ，
// 因此此函数被认为广泛在基本数据类型的切片上使用
func SliceByFunc[T any, S ~*[]T](
	operator IO, x S,
	marshal_func func(x *T),
) {
	len := len(*x)
	if len > MaxLenSlice {
		panic(fmt.Sprintf("SliceByFunc: The length of the target slice is out of the max limited %d; len = %d", MaxLenSlice, len))
	}
	// pre check
	length := uint32(len)
	operator.Uint32(&length)
	// wrire or read length
	_, success := operator.(*Reader)
	if success {
		*x = make([]T, length)
	}
	// make slice if meet read mode
	for i := uint32(0); i < length; i++ {
		marshal_func(&(*x)[i])
	}
	// marshal or unmarshal data included
}

/*
从映射 x(map[K]V) 读取或向映射 x(map[K]V) 写入数据，
并以 IO 操作流 operator 为媒介编码或解码数据。

M 指代封装有在 V 上实现了基于 IO 操作流的 解码/编码 接口，
这被用于映射中值的 解码/编码 ；
key_marshal_func 指代基于 IO 操作流而实现的函数，
被用于映射中键的 解码/编码 。

key_marshal_func 与 M 的不同之处在于，
前者要求直接提供对应的函数用于键的 解码/编码 ，
因此映射的键在通常情况下应该是基本数据类型
*/
func Map[K comparable, V any, M Marshal[V]](
	operator IO, x *map[K]V,
	key_marshal_func func(x *K),
) {
	len := len(*x)
	if len > MaxLenMap {
		panic(fmt.Sprintf("Map: The length of the target map is out of the max limited %d; len = %d", MaxLenSlice, len))
	}
	// pre check
	length := uint16(len)
	operator.Uint16(&length)
	// wrire or read length
	switch operator.(type) {
	case *Writer:
		for key, value := range *x {
			key_marshal_func(&key)
			M(&value).Marshal(operator)
		}
		// write data
	case *Reader:
		*x = make(map[K]V)
		// make map
		for i := uint16(0); i < length; i++ {
			var key K
			var value V
			// prepare
			key_marshal_func(&key)
			M(&value).Marshal(operator)
			// read key and data
			(*x)[key] = value
			// set data
		}
		// read data
	}
	// write or read data
}

/*
从映射 x(map[K]V) 读取或向映射 x(map[K]V) 写入数据，
并以 IO 操作流 operator 为媒介编码或解码数据。

key_marshal_func 和 value_marshal_func 均指代基于
IO 操作流而实现函数。前者被用于映射中键的 解码/编码 ，
而后者则被用于映射中值的 解码/编码 。

此函数与 Map 的不同之处在于，
此处要求直接提供对应的函数用于映射的 解码/编码 ，
因此此函数被认为广泛在键和值均为基本数据类型的映射上使用
*/
func MapByFunc[K comparable, V any](
	operator IO, x *map[K]V,
	key_marshal_func func(x *K),
	value_marshal_func func(x *V),
) {
	len := len(*x)
	if len > MaxLenMap {
		panic(fmt.Sprintf("MapByFunc: The length of the target map is out of the max limited %d; len = %d", MaxLenSlice, len))
	}
	// pre check
	length := uint16(len)
	operator.Uint16(&length)
	// wrire or read length
	switch operator.(type) {
	case *Writer:
		for key, value := range *x {
			key_marshal_func(&key)
			value_marshal_func(&value)
		}
		// write data
	case *Reader:
		*x = make(map[K]V)
		// make map
		for i := uint16(0); i < length; i++ {
			var key K
			var value V
			// prepare
			key_marshal_func(&key)
			value_marshal_func(&value)
			// read key and data
			(*x)[key] = value
			// set data
		}
		// read data
	}
	// write or read data
}
