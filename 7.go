package main

import "fmt"

func main() {
	var n = 100;
	fmt.Printf("%T\n",n ) // 查看类型
	fmt.Printf("%v\n",n )//查看值
	fmt.Printf("%b\n",n)//查看二进制
	fmt.Printf("%d\n",n )//查看正数
	fmt.Printf("%o\n", n)//查看八进制数
	fmt.Printf("%x\n", n)//查看十六进制数
	var s = " Hello Golang"
	fmt.Printf("字符串%s\n",s )
	fmt.Printf("字符串%v\n",s)
	fmt.Printf("字符串%#v\n",s)
}
