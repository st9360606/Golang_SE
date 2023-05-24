package main

// import "fmt"

// func f1() {
// 	//add
// 	var s1 = []int{}
// 	s1 = append(s1, 100)
// 	s1 = append(s1, 200)
// 	s1 = append(s1, 300)
// 	fmt.Printf("s1: %v\n", s1)

// 	//delete
// 	var s2 = []int{1, 2, 3, 4}
// 	//刪除2
// 	s2 = append(s2[:2], s2[3:]...)
// 	fmt.Printf("s2: %v\n", s2)

// 	//update
// 	s2[0] = 2
// 	fmt.Printf("s2: %v\n", s2)
// }

// func f2() {
// 	//query
// 	var s2 = []int{1, 2, 3, 4}
// 	var key = 2
// 	for i, v := range s2 {
// 		if v == key {
// 			fmt.Printf("i: %v\n", i)
// 		}

// 	}

// 	//copy的操作
// 	var s1 = s2 //這樣是址傳遞
// 	s2[0] = 10
// 	fmt.Printf("s1: %v\n", s1) //s1: [100 2 3 4]
// 	fmt.Printf("s2: %v\n", s2) //s2: [100 2 3 4]

// 	var s4 = []int{1, 2, 3, 4}
// 	var s3 = make([]int, 4)
// 	copy(s3, s4) //這樣是值傳遞
// 	s3[0] = 500
// 	fmt.Printf("s3: %v\n", s3) //s3: [100 2 3 4]
// 	fmt.Printf("s4: %v\n", s4) //s2: [100 2 3 4]

// }

// func main() {
// 	// f1()
// 	f2()
// }
