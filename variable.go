package main

import "fmt"

func main(){
	fmt.Println("muhaha")
	var toto string
	var ff string
	var xx int
	var jj string

	toto="toto"
	ff="xxx"
	fmt.Println(toto+ff)
	fmt.Println(xx)
	fmt.Println(jj)
	var j=false
	fmt.Println(j)
	noVar :=false
	fmt.Println(noVar)
a,b,c:=1,2,3
fmt.Println(a+b+c)
var (
	e int
	g bool=true
)
fmt.Println(g)
fmt.Println(e)
println(a,b,c,e,g)
println(&b)
h:=999
h=888
println(h)
_,c=2,33333//_用来抛弃不需要的数值
println(c)
const aa, bb, cc = 1, false, "uuu"
fmt.Println(aa,bb,cc)

}