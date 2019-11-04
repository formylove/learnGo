package main

import "fmt"

func main() {
 var ptr *int

 fmt.Printf("ptr 的值为 : %x\n", ptr )
 
 if(ptr == nil){
 println("is null")
 }
}