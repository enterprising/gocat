package main

import "fmt"

// 函数闭包
// nextNum 这个函数返回了一个匿名函数，这个匿名函数记住了nextNum中i+j的值，并改变的i,j的值。于是形成了一个闭包的用法
func main() {
	nextNumFunc := nextNum()
	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}
}
func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}
