package main
import "fmt"
func foo(num...int)int{
	sum :=0
	for _, element := range num{
		sum=sum+element

	}
	return sum
}
func bar(num[]int)int{
	sum :=0
	for _,v :=range num{
		sum=sum+v
	}
	return sum
}
func main(){
	xs :=[]int{1,2,3,4,5}
	fmt.Println(foo(xs...))
	fmt.Println(bar(xs))
}

