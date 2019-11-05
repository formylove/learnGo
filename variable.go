package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3 //初始化也可以使用expand
	var (
		e int
		g bool = true
	)
	println(a, b, c, e, g)
	println(&b)
	h := 999
	h = 888
	println(h)
	_, c = 2, 33333 //_用来抛弃不需要的数值
	println(c)
	const aa, bb, cc = 1, false, "uuu"//可以同时设置多个类型常量
	fmt.Println(aa, bb, cc)

}
