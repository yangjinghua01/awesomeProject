package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("\ns1=%v,len(s1)=%d,cap(s1)=%d", s2, len(s2), cap(s2))
}
