package main
import "fmt"
func main(){
var arr=make([]int,5,10)


arr[1]=99
// arr[6]=199
fmt.Println(arr,cap(arr),len(arr))


arr2 := []int{2, 3, 5, 7, 11, 13}
sli := arr2[1:4]
fmt.Println(sli)
fmt.Println("len of sli: ",len(sli))
//在数组中由于长度固定不可变，因此len(arr)和cap(arr)的输出永远相同
fmt.Println("len of arr2: ",len(arr2))
fmt.Println("cap of arr2: ",cap(arr2))
fmt.Println("cap of sli: ",cap(sli))
}