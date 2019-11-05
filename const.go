package main

import "fmt"
const (
aa = "iota"  //iota
bb     //iota
cc    //iota

)
const (
i=1<<iota
z=1<<iota
j=3<<iota
k
l
)
/* j=3：左移 1 位,变为二进制 110, 即 6;
k=3：左移 2 位,变为二进制 1100, 即 12;
l=3：左移 3 位,变为二进制 11000,即 24。 */
func main() {
const (
a = iota  //0
b     //1
c     //2
d = "ha"  //独立值，iota += 1
e     //"ha"  iota += 1
f = 100  //iota +=1
g     //100 iota +=1
h = iota  //7,恢复计数
i     //8
)
fmt.Println(a,b,c)
fmt.Println(d,e,f,g,h,i)
fmt.Println(aa,bb,cc)
fmt.Println(i,j,k,l,z)
}