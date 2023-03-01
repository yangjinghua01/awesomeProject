package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		println("成年人")
	} else if age >= 16 {
		println("半成年人")
	} else {
		println("小孩子")
	}
	for i := 0; i < 10; i++ {
		println(i)
	}
	s := "asdfghjkl"
	for i, i2 := range s {
		fmt.Printf("%d", i)
		fmt.Printf("%c ", i2)
	}
	//	哑元变量，不想用到的都直接扔给它
	for _, v := range s {
		fmt.Printf("遍历获得的值%c", v)
	}
	//	9*9表
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("%d*%d=%d\t ", j, i, i*j)
		}
		println()
	}
}
