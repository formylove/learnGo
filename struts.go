package main

import "fmt"

type Books struct{
title string
author string
subject string
book_id int
}


func main(){

//创建一个新的结构体
fmt.Println(Books{"Go语言","www.runoob.com","Go语言教程",6495407})

//也可以使用key=>value格式
fmt.Println(Books{title:"Go语言",author:"www.runoob.com",subject:"Go语言教程",book_id:6495407})

//忽略的字段为0或空
fmt.Println(Books{title:"Go语言",author:"www.runoob.com"})
}