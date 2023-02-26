package main

import (
	"fmt"
	"strings"
)

//字符串
func main() {
	// \本来是具有特殊含义的，我应该告诉程序我写的\就是一个单纯的\
	var path = "\"D\\Go\\src\\codeyang\\study\\note\""
	fmt.Printf("%s", path)
	//	多行字符串
	s2 := `
世人熙熙皆为利来
世人攘攘皆为利往
`
	fmt.Printf(s2)
	fmt.Printf("%d", len(s2))
	// 字符串拼接
	firstName := "敬华"
	lastName := "杨"
	fmt.Printf("\n%s%s", lastName,firstName)
	allName :=lastName+firstName
	println("\n",allName)
	ss1 := fmt.Sprintf("%s%s", lastName, firstName)
	fmt.Printf("%s", ss1)
	//判断一个字符串里是否包含另一个字符串和java的一致
	println(strings.Contains(ss1,firstName))
	//判断字符串前缀是否包含
	println(strings.HasPrefix(ss1,lastName))
	//判断字符串后缀是否包含
	println(strings.HasSuffix(ss1,lastName))
	s4:= "abcdeb"
	println(strings.Index(s4,"c"))//判断子串第一次出现的位置
	println(strings.LastIndex(s4,"b"))//判断子串最后一次出现的位置
	s5:="杨敬华杨敬"
	println(strings.Index(s5,"敬"))//判断子串第一次出现的位置
	println(strings.LastIndex(s5,"敬"))//判断子串最后一次出现的位置
}
