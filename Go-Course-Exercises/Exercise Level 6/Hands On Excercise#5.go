package main
import (
	"fmt"
)
type shape interface{
	area() float64
}
type circle struct{
	r float64
}
type rect struct{
	width float64
	height float64
}
func(c circle)area()float64{
	return 3.14*c.r*c.r
}
func(r rect) area()float64{
	return r.width*r.height
}
func info(sh shape){
	fmt.Println(sh.area())
}

func main(){
c1 :=circle{r:2.3}
r1 :=rect{width :2, height:3}
info(c1)
info(r1)
}