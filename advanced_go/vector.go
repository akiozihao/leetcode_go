package main

import "fmt"

func main() {
	var a1 = [...]int{1, 2, 3}
	var b = &a1

	fmt.Println(a1[0], a1[1])
	fmt.Println(b[0], b[1])

	for i, v := range b {
		fmt.Println(i, v)
	}

	for i := range a1 {
		fmt.Printf("a[%d]: %d\n", i, a1[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d", i, v)
	}
	for i := 0; i < len(b); i++ {
		fmt.Printf("b[%d]: %d\n", i, b[i])
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

	var s1 = [2]string{"你好", "世界"}
	for i := range s1 {
		fmt.Println(s1[i])
	}
	fmt.Println("\x2c\x20")
	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界


}
