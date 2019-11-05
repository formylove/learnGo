package main

import "fmt"

func main() {
 var ptr *int

 fmt.Printf("ptr 的值为 : %x\n", ptr )
 
 if(ptr == nil){//判断是否是空指针︰括号可以加可以不加
 println("is null")
 }
}