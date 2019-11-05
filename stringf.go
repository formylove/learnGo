package main
import "fmt"
func main(){
strFormat := `
  Cannot proceed, the divider is zero.
  dividee: %d
  divider: 0
`

x:=fmt.Sprintf(strFormat,22222)

fmt.Println(x)
}