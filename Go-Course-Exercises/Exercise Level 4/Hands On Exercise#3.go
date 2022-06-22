package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice[:5]) //0,1,2,3,4,5
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])

	slice1 := slice[1:6] //1,2,3,4,5
	fmt.Println(slice1)
}
