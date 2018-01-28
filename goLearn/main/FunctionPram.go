package main

import "fmt"

// 函数不定参数

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 传参数
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
func sum(nums ...int) {
	fmt.Println(nums)

	total := 0
	for _, num := range nums {
		//要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}
