package main

import "fmt"

// 定义常量和变量
func main() {
	// 声明初始化一个变量
	var x int = 100;
	var str string = "hello world"

	// 声明初始化多个变量
	var i, j, k int = 1, 2, 3;

	// 不用指明类型，通过初始化来推导
	var b = true;

	y := 100

	// 定义常量
	const s string = "hello go"
	const pi float32 = 3.1415926

	fmt.Println(x)
	fmt.Println(str)
	fmt.Println(i, j, k)
	fmt.Println(b)
	fmt.Println(y)
	fmt.Println("-------")
	fmt.Println(s)
	fmt.Println(pi)
}
