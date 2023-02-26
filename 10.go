package main

import "fmt"

func main() {
	age:=18
	if age>=18 {
		println("成年人")
	}else if age>=16 {
		println("半成年人")
	}else {
		println("小孩子")
	}
	for i := 0; i < 10; i++ {
		println(i)
	}
	s:="asdfghjkl"
	for i, i2 := range s {
		fmt.Printf("%d", i)
		fmt.Printf("%c", i2)
	}
}
