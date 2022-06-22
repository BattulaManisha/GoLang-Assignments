package main
import "fmt"
func main(){
	f :=func(){ fmt.Println("Assign func to variable")}
	f()
	fmt.Printf("f is type of %T",f)
}