package main
import "fmt"


func main(){
	const(
		Apple int =iota// It is used to increment the constants
		orange
		mango
		papaya
		)
	
	fmt.Printf("Apple : %v\nOrange : %v\nMango : %v\nPapaya : %v",Apple,orange,mango,papaya)
}