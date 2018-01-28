package main

import "fmt"

//函数返回多个值。一个正常值一个错误
func main() {
	v, e := multi_ret("one")
	fmt.Println(v, e)

	v, e = multi_ret("four")
	fmt.Println(v, e)

	// 通常的用法
	if v, e = multi_ret("four"); e {
		println("正常返回")
	} else {
		println("出错返回")
	}

}
func multi_ret(key string) (int, bool) {
	m := map[string]int{"one":1, "two":2, "three":3}

	var err bool
	var val int

	val, err = m[key]

	return val, err
}