package main

import "fmt"

// 函数

func main() {
	var i int
	i = max(4, 5)
	fmt.Println(i)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
