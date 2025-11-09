package main
import "fmt"

func main(){
	var num int = 42
	fmt.Printf("The number is: %T\n", num)	
	var str string = "Hello, Go!"
	fmt.Printf("The string is: %T\n", str)

	s := ConvertedString(num)
	fmt.Printf("Converted string: %s, Type: %T\n", s, s)
}