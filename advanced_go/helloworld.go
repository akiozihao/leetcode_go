package main 

import "fmt"
// 比较特殊的是，Go语言闭包函数对外部变量是以引用的方式使用
func main1() {
	fmt.Println("你好，世界")
}