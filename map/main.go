package main

import "fmt"

func main(){
	fmt.Println("path package in golang")
	students := make(map[string]int);

	students["manoj"] = 100
	students["ram"] = 200
	students["shyam"] = 300
	for index , value := range students{
		fmt.Printf("key is %s and value is %d\n",index , value)
	}
	fmt.Println(students["manoj"])
}