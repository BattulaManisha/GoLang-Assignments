package main
import "fmt"
func foo() int{
	return 0
}
func bar()(int ,string ){
	return 10,"manu"

}
func main(){
	fmt.Println(foo())
	fmt.Println(bar())
}