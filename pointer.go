package main

import "fmt"

func main() {
 var a int = 4
 var b int32
 var c float32
 var ptr *int

 /* 运算符实例 */
 fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
 fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
 fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

 /* & 和 * 运算符实例 */
 ptr = &a   /* 'ptr' 包含了 'a' 变量的地址 */
 fmt.Printf("a 的值为 %d\n", a);
 fmt.Printf("*ptr 为 %d\n", *ptr);//指向a


var d int= 20  /* 声明实际变量 */
 var ip *int    /* 声明指针变量 */

 ip = &d /* 指针变量的存储地址 */

 fmt.Printf("d 变量的地址是: %x\n", &d )

 /* 指针变量的存储地址 */
 fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

 /* 使用指针访问值 */
 fmt.Printf("*ip 变量的值: %d\n", *ip )
}