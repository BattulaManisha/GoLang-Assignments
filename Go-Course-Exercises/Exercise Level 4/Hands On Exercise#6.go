package main
import "fmt"
func main(){
	
	slice := make([]string, 28)
	slice = []string{
		"Alabama","Alaska","Arizona","arkansas","California","Colorada","connecticut","Delaware","Florida","Georgia","Hawaii","Illinois","Indiana","Lowa","Kansas","Kentucky","Louisiana","Maine","Maryland","Massachusetts","Michigan","Minnesota","Mississipi","Missouri","Montana","Nebraska","Nevada","New Hampshire","New Jersey","New Mexico","New York","North Carolina","North Dakota","Ohio","Oklahoma","Oregon","Pennsylvania","Rhode Island","South Carolina","South Dakota","Tennsessee","Texas","Utah","Vermont","Virginia","Washington","West Virginia","Wisconsin","Wyoming"}
	for i :=0;i<len(slice);i++ {
		fmt.Printf("%d: %s\n",i+1,slice[i])
	}
}