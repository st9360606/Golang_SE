package main

import "fmt"

func f1() {
	var v1 = map[string]string{"name": "kuku", "age": "20"}
	for k := range v1 {
		fmt.Printf("k: %v\n", k)
	}
	
	for k, v := range v1 {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func main() {
	f1()
}
