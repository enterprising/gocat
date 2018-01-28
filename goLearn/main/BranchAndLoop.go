package main

import "fmt"

//分支循环语句
func main() {
	var x int = 10
	// if
	if x % 2 == 0 {
		fmt.Println("x%2=0")
	}

	x = 9
	// if-else
	if x % 2 == 0 {
		fmt.Println("x%2=0")
	} else {
		fmt.Println("x%2!=0")
	}

	// if - else if - else
	if x == 0 {
		fmt.Println("x=0")
	} else if x == 2 {
		fmt.Println("x==2")
	} else {
		fmt.Println("x=", x)
	}

	// switch
	var i int = 6
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 6:
		fmt.Println("six")
	case 7, 8, 9, 10:
		fmt.Println("xxxx")
	default:
		fmt.Println("invalid value!")
	}

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("===========")

	// 精简的for语句
	j := 1
	for j < 10 {
		fmt.Println(j)
		j++
	}
	fmt.Println("===========")

	// 死循环的for语句，相当于for ; ;
	k := 1
	for {
		if k > 10 {
			fmt.Println(k)
			break
		}
		k++
	}

}
