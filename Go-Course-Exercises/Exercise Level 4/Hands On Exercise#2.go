package main
import "fmt"
func main(){
	slice :=[]int{1,4,6,7,9,10,3,7,9,2}
	for _, element := range slice {
		fmt.Printf("%v\n",element)
	} 
	fmt.Printf("%T",slice)
}