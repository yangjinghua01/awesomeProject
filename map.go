package main

import "fmt"

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["yjh"] = 18
	m1["xxgc"] = 17
	fmt.Println(m1)
	fmt.Println(m1["yjh"])
	k, v := m1["xxx"]
	if v {
		fmt.Println(v, k)
	} else {
		fmt.Println(v, k)
	}
	for _, v := range m1 {
		fmt.Println(v)
	}
	//删除
	delete(m1, "yjh")
	fmt.Println(m1)
	delete(m1, "caonima")
	fmt.Println(m1)
}
