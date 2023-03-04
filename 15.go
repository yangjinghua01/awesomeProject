package main

import "fmt"

// 多维数组
func main() {
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{1, 2},
		[2]int{1, 2},
	}
	fmt.Println(a11)
	for a := 0; a < len(a11); a++ {
		fmt.Println(a11[a])
		for i := 0; i < len(a11[i]); i++ {
			fmt.Println(a11[a][i])
		}
	}
	//这是值类型传递类似深拷贝
	b1 := [3]int{10000, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1[0])
	//求数组的和【1,3,5,7,8】
	test := [5]int{1, 3, 5, 7, 8}
	var sum int
	for i := 0; i < len(test); i++ {
		sum = test[i] + sum
	}
	println(sum)
	sums := [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(sums); i++ {
		for j := 0; j < len(sums); j++ {
			if j+1 < len(sums) {
				if sums[i]+sums[j+1] == 8 {
					fmt.Println(i, j+1)
				}
			}

		}
	}
}
