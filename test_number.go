package main

// import (
// 	"fmt"
// 	"math"
// 	"unsafe"
// )

// func main() {
// 	var u uint     // 8 Bytes = 64 Bits on 64 Bit OS
// 	var u8 uint8   // 1 Byte = 8 Bits
// 	var u16 uint16 // 2 Bytes = 16 Bits
// 	var u32 uint32 // 4 Bytes = 32 Bits
// 	var u64 uint64 // 8 Bytes = 64 Bits

// 	var i int     // 8 Bytes = 64 Bits on 64 Bit OS
// 	var i8 int8   // 1 Byte = 8 Bits
// 	var i16 int16 // 2 Bytes = 16 Bits
// 	var i32 int32 // 4 Bytes = 32 Bits
// 	var i64 int64 // 8 Bytes = 64 Bits

// 	var f32 float32
// 	var f64 float64

// 	fmt.Printf("u: %T, %d\n", u, unsafe.Sizeof(u))
// 	fmt.Printf("u8: %T, %d\n", u8, unsafe.Sizeof(u8))
// 	fmt.Printf("u16: %T, %d\n", u16, unsafe.Sizeof(u16))
// 	fmt.Printf("u32: %T, %d\n", u32, unsafe.Sizeof(u32))
// 	fmt.Printf("u64: %T, %d\n", u64, unsafe.Sizeof(u64))
// 	fmt.Printf("----------------------------------------\n")
// 	fmt.Printf("i: %T, %d\n", i, unsafe.Sizeof(i))
// 	fmt.Printf("i8: %T, %d\n", i8, unsafe.Sizeof(i8))
// 	fmt.Printf("i16: %T, %d\n", i16, unsafe.Sizeof(i16))
// 	fmt.Printf("i32: %T, %d\n", i32, unsafe.Sizeof(i32))
// 	fmt.Printf("i64: %T, %d\n", i64, unsafe.Sizeof(i64))
// 	fmt.Printf("----------------------------------------\n")
// 	fmt.Printf("f32: %T, %d\n", f32, unsafe.Sizeof(f32))
// 	fmt.Printf("f64: %T, %d\n", f64, unsafe.Sizeof(f64))

// 	fmt.Printf("math.Pi: %f\n", math.Pi)
// 	fmt.Printf("math.Pi: %.2f\n", math.Pi)
// }