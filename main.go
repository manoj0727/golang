package main

import(
	"fmt"
	"mylearning/myutil"
) 
func main(){
	fmt.Println("master go lang")
	fmt.Println("---------------")

	myutil.PrintMessage("HELLO BHAI")

	var name string = "MANOJ"
	var name1 int = 100
	fmt.Println(name1)
	var version = "version latest"
	fmt.Println(version)
	fmt.Println(name)
	fmt.Println(43);

	person :="MANOJ KUMAWAT"
	fmt.Println(person)

	var Public = "data ki external use happen exported"
	var private = "data ki internal unexported"

	fmt.Println(Public);
	fmt.Println(private);
}