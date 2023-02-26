package main

//单独声明常量
const pi = 3.1415926

//批量声明变量
const (
	statusOk = 200
	notFount = 404
)

//批量声明常量值时，如果某一行声明后没有赋值，默认就和上一行一致。
const (
	n1 = 100
	n2
	n3 = 110
	n4
)
const (
	a1 = iota
	a2 = iota
	a3 = iota
)
//定义数量级
const(
	_ = iota
	KB = 1<< (10 *iota)
	MB = 1<< (10 *iota)
	GB = 1<< (10 *iota)
	TB = 1<< (10 *iota)
	PB = 1<< (10 *iota)
)
func main() {
	println(n1)
	println(n2)
	println(n3)
	println(n4)
	println(a1)
	println(a2)
	println(a3)
}
