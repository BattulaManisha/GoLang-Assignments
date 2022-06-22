package main
import "fmt"
func main(){
	 arr :=[5]int{1,5,8,4,9}
	for _, element := range arr {
		fmt.Printf("%v\n",element)
	}
	fmt.Printf("%T",arr)

}