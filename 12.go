package main

// goto
func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto xx
		}
		println(i)
	}
xx:
	println("goto出来执行")
}
