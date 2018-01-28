package main

import "fmt"

// 数组
func main() {
	var a [5]int;
	fmt.Println("array a:", a)

	a[1] = 10
	a[3] = 30

	fmt.Println("assign:", a)
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init b:", b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2d: ", c)

	// 数组的切片，和Python一样
	fmt.Println("----------")
	array := [5]int{1, 2, 3, 4, 5}
	tempArray := array[2:4]  // a[2]和a[3]，但不包括a[4]
	fmt.Println(tempArray)

	tempArray = array[:4]  //从 a[0] 到 a[4]，但不包括a[4]
	fmt.Println(tempArray)

	tempArray = array[2:] //从a[2]到a[4]，但不包括a[2]
	fmt.Println(tempArray)

}