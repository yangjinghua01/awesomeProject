package main

import "fmt"

// copy
func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	a3 := make([]int, 5, 5)
	copy(a3, a1) //copy
	fmt.Printf("a3=%v,a2=%v,a1=%v", a3, a2, a1)
	a1[0] = 10000
	fmt.Printf("\na3=%v,a2=%v,a1=%v", a3, a2, a1)
}
