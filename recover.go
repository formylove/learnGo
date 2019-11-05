package main
import "fmt"
func main(){
    defer func(){
        if err := recover() ; err != nil {
            fmt.Println(err)
        }
    }()
    defer func(){
        panic("three")
    }()
    defer func(){
        panic("two")
    }()
    panic("one")
}