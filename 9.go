package main

import "fmt"

func main() {
	s2:="脑瓜疼"
	s2 ="头发疼"
	fmt.Printf("%s", s2)
	//字符串修改
	s3 :=[]rune(s2) // 这里相当于把脑瓜疼这个字符串转化为['脑','瓜','疼']
	s3[0] = '西'//应为是字符而不是字符串用''
	fmt.Printf("%s", string(s3))
}
