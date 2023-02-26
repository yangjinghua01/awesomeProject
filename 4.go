package main

import "fmt"

//整形

func main() {
	i1 :=101
	fmt.Printf("%d\n",i1)
	fmt.Printf("转八进制%o\n", i1) //把十进制转成八进制
	fmt.Printf("转十六进制%x\n", i1)//把十进制转十六进制
	fmt.Printf("转二进制%b\n", i1)//把十进制转二进制
	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制的数
	i3 :=0x1234567
	fmt.Printf("%d\n",i3)
	//查看变量的类型
	fmt.Printf("%T",i3)
	//声明int8类型的变量
	i4:= int8(9)
	fmt.Printf("i4=%d\n",i4 )
	fmt.Printf("%T\n",i4 )
}
