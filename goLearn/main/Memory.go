package main

import "fmt"

/**
内存分配：
new和make
new(T): 将内存清零，而不是初始化内存。 返回了一个指向新分配的类型为T的零值的指针
make: 返回一个被初始化了的实例
eg:
make([]int,10,100)  分配了一个整型数组，长度为10，容量100，并返回前10个数组的切片
 */
func main() {

	var p *[]int = new([]int) //为切片结构分配内存； *p == nil 很少使用
	var v []int = make([]int, 10) //切片v现在是对一个新的有10个整数的数组的引用

	fmt.Println(p) //输出: &[]

	*p = make([]int, 10, 10)
	fmt.Println(p)
	fmt.Println((*p)[2])

	// 习惯用法
	v = make([]int, 10)
	fmt.Println(v)

}
