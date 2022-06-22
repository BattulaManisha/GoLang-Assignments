package main
import "fmt"
type person struct{
	first string
	last string
	age int
}
func changeMe(p *person){
	p.first="Manisha"
	p.last="Battula"
	p.age=23
}
func main(){
	p1:=person{
		first :"Mounika",
		last :"Gosu",
		age :20,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)


	

}