package main
import "fmt"
func main(){
	x :=4
	fmt.Printf("Decimal: %d\nBinary: %b\nHexaDecimal: %#X",x,x,x)
	y :=x<<1
	fmt.Printf("\nDecimal: %d\nBinary: %b\nHexaDecimal: %#X",y,y,y)

}