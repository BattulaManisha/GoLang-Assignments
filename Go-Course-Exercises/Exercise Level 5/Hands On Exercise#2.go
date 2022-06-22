package main

import "fmt"

func main() {
	type Person struct {
		firstName, lastName string
		favIceCreamFlavour  []string
	}

	me := Person{
		firstName:          "Manisha",
		lastName:           "Battula",
		favIceCreamFlavour: []string{"Vennala", "Blueberry"},
	}
	bob := Person{
		firstName:          string("Raju"),
		lastName:           string("Gosu"),
		favIceCreamFlavour: []string{"Chocolate"},
	}

	people := map[string]Person{
		me.lastName:  me,
		bob.lastName: bob,
	}

	for _, v := range people {
		fmt.Printf("name: %v %v\nfavorite ice cream flavors: ", v.firstName, v.lastName)
		for _, v := range v.favIceCreamFlavour {
			fmt.Printf("%v ", v)
		}
		fmt.Println("\n")
	}
}