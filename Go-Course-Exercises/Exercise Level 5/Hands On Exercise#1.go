package main
import "fmt"
type person struct{
	firstname string
	lastname string
	icecream []string
}
func main(){
	p1 :=person{
		firstname : "Manisha",
		lastname :"Battula",
		icecream: []string{"vennala","Blueberry"},
	}
	p2 :=person{
		firstname :"Ravi",
		lastname :"Gosu",
		icecream: []string{"chocolate"},
	}
	fmt.Printf("name: %v %v\nfavorite ice cream flavors: ", p1.firstname, p1.lastname)
	for _, v := range p1.icecream {
		fmt.Printf("%v ", v)
	}
	fmt.Println("\n")

	fmt.Printf("name: %v %v\nfavorite ice cream flavors: ", p2.firstname, p2.lastname)
	for _, v := range p2.icecream {
		fmt.Printf("%v ", v)
	}
	
	


	}
