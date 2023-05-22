package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	fmt.Printf("s: %v\n", s)

	s1 := `
	line 1
	line 2
	line 3
	`
	fmt.Printf("s1: %v\n", s1)

	//字符串聯接
	s2 := "tom"
	s3 := "20"
	msg := s2 + s3
	fmt.Printf("msg: %v\n", msg)

	msg1 := fmt.Sprintf("name=%s , age=%s", s2, s3)
	fmt.Printf("msg1: %v\n", msg1)

	msg2 := strings.Join([]string{s2, s3}, ",")
	fmt.Printf("msg2: %v\n", msg2)

	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("30")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	//轉譯字符
	url := "c:\\program files\\a.txt"
	fmt.Printf("url: %v\n", url)

	//字符串切片動作
	a := "hello world"
	b := 2
	c := 5
	fmt.Printf("a[b]: %c\n", a[b])
	fmt.Printf("a[b:c]: %v\n", a[b:c])
	fmt.Printf("a[b:]: %v\n", a[b:])
	fmt.Printf("a[:c]: %v\n", a[:c])

	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("strings.Split(s,\"\"):%v\n", strings.Split(s, " "))
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))
	fmt.Printf("strings.HasPrefix(\"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("strings.HasSuffix(\"world\"): %v\n", strings.HasSuffix(s, "world"))
	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))
}
