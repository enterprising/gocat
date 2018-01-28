package main

import "fmt"

func main() {
	m := make(map[string]int) // 使用make创建一个空的map

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)  //不能保证顺序输出
	fmt.Println(len(m))

	v := m["two"]
	fmt.Println(v)

	delete(m, "two")
	fmt.Println(m)

	m1 := map[string]int{"one":1, "two":2, "three":3}
	fmt.Println(m1)

	for key, value := range m1 {
		fmt.Printf("%s => %d\n", key, value)
	}
}
