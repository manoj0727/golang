package main

import "fmt"

func SimpleFunc(){
	fmt.Println("this is simple function")
}
func add(a int, b int) int{
	return a+b
}
func multiply(a int, b int) int{
	return a*b
}
func main(){
	fmt.Println("we are in master")
	SimpleFunc()

	ans := add(4,5)
	fmt.Println("addition is:", ans)


	fmt.Println(multiply(4,5))
}