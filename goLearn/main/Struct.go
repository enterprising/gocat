package main

import "fmt"

// 结构体

type rect struct {
	width, heigh int
}

func (r *rect) perimeter() int {
	return 2 * (r.heigh + r.width)
}

func (r *rect) area() int {
	return r.width * r.heigh
}

func main() {
	r := rect{width:10, heigh:15}

	fmt.Println("周长：", r.perimeter())
	fmt.Println("面积：", r.area())

	rp := &r
	fmt.Println("周长：", rp.perimeter())
	fmt.Println("面积：", rp.area())
}
