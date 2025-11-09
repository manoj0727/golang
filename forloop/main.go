package main

import "fmt"

func main() {
	fmt.Println("for loop in golang")

	for   i:=0;i<7;i++{
		fmt.Println("value of i is:",i)
	}
	var count = 5;
	for{
		fmt.Println("infinite loop")
		count--;
		if count ==0{
			break;
		}
	}
}