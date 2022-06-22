package main

import "fmt"

func main() {
	a := true == true
	b := 99 <= 67
	c := 33 >= 33
	x := "Manisha Battula" != "Manisha Battula"
	y := -9 < 22
	z := 100-37 > 6*6

	fmt.Printf("a = %v\nb = %v\nc = %v\nx = %v\ny = %v\nz = %v\n", a, b, c, x, y, z)
}