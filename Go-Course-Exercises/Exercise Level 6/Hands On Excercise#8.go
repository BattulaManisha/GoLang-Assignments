package main
import "fmt"
func foo() func(s string) string{
	return func(s string) string {
		return "hello "+s
	}

}
func main(){
	v:=foo()
	fmt.Println(v("Manisha Battula"))
}