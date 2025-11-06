package main
import "fmt"

func divide(a ,b float64) (float64,error){
	if b==0{
		return 0, fmt.Errorf("division by zero error")
	}
	 return a/b, nil
}
func main() {
	fmt.Println("This is error by system")
	ans, _ := divide(10,0);
	fmt.Println("The division is:", ans)
}