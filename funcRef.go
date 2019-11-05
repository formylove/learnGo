package main

import (
 "fmt"
 "math"
)
// 声明一个函数类型
type cb func(int) int

func main(){
 /* 声明函数变量 */
 getSquareRoot := func(x float64) float64 {
   return math.Sqrt(x)
 }

 /* 使用函数 */
 fmt.Println(getSquareRoot(9))

testCallBack(9, callBack)

 testCallBack(2, func(x int) int {
        fmt.Printf("我是回调，x：%d\n", x)
        return x
    })
     
}
func testCallBack(x int, f cb) {
    f(x)
}

func callBack(x int) int {
    fmt.Printf("我是回调，x：%d\n", x)
    return x
}