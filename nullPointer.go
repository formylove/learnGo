package main

import "fmt"

func main() {
 var ptr *int

 fmt.Printf("ptr 的值为 : %x\n", ptr )
 
 if(ptr == nil){//判断是否是空指针
 println("is null")
 }
}