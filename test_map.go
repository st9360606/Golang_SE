package main

// import "fmt"

// func f1() {
// 	//類型的聲明
// 	var m1 map[string]string
// 	fmt.Printf("m1: %v\n", m1)
// 	fmt.Printf("m1: %T\n", m1)

// 	m2 := make(map[string]string)
// 	fmt.Printf("m2: %v\n", m2)
// 	fmt.Printf("m2: %T\n", m2)
// }

// func f2() {
// 	//兩種初始化方式
// 	var v1 = map[string]string{"name": "kuku", "age": "20"}
// 	fmt.Printf("v1: %v\n", v1)

// 	v2 := make(map[string]string)
// 	v2["name"] = "opop"
// 	v2["age"] = "90"
// 	fmt.Printf("v2: %v\n", v2)
// }

// func f3() {
// 	var v1 = map[string]string{"name": "kuku", "age": "20"}
// 	var key = "name"
// 	var v = v1[key]
// 	fmt.Printf("v: %v\n", v)
// }

// func f4() {
// 	//golang中判斷某個鍵是否存在的特殊寫法
// 	var v1 = map[string]string{"name": "kuku", "age": "20"}
// 	var k1 = "name"
// 	var k2 = "lala"
// 	v, ok := v1[k1]
// 	fmt.Printf("v: %v\n", v)
// 	fmt.Printf("ok: %v\n", ok)
// 	fmt.Println("============================")
// 	v, ok = v1[k2]
// 	fmt.Printf("v: %v\n", v)
// 	fmt.Printf("ok: %v\n", ok)
// }

// func main() {
// 	// f1()
// 	// f2()
// 	// f3()
// 	f4()
// }
