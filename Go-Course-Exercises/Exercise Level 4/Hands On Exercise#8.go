package main

import "fmt"

func main() {
	myMap := map[string][]string{
		"bond-james":           {"Shaken", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":   {"Being evil", "Ice cream", "Sunsets"},
	}

	for k, slice := range myMap {
		fmt.Println(k)
		for i, v := range slice {
			fmt.Printf("%d: %v\n", i+1, v)
		}
		fmt.Println()
	}
}