package main

import "fmt"

//声明变量(单个声明)
var name string
var age int
var isOk bool

//批量声明
var (
	name1 string //""
	age1  int    //0
	isOk1 bool   //false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//Go语言中变量声明必须使用，不实用就编译不过去
	print(isOk)                //在终端中输出要打印的内容
	fmt.Printf("name%s", name) //%s:占位符  使用name这个变量的值去替代占位符
	fmt.Println(name)          //打印完指定的内容之后会在后面加一个换行符
	//简短变量声明
	s3:="hahahh"//这种方式只能在函数里使用。
	print(s3)//同一个作用域中不能重复声明同名的变量
}
