package main

import "fmt"

func main() {
	slice2D := [][]string{{"James", "Bond"}, {"Shaken", "Not stirred"}, {"Miss", "Moneypenny"},{"Hellooooo","James"}}

	for i := range slice2D {
		fmt.Printf("slice %d:\n", i)
		for _, v2 := range slice2D[i] {
			fmt.Printf("%s ", v2)
		}
		fmt.Println("\n")
	}
}