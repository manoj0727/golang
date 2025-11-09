package main

import "fmt"

type Person struct{
	Firstname string
	lastname string
	age int
}

func main(){
		var prince Person
		fmt.Println("Prince Person", prince)
		prince.Firstname ="Manoj"
		prince.lastname ="Kumawat"
		prince.age = 24
		fmt.Println("Prince Person", prince)

		person1 := Person{
			Firstname: "Ram",
			lastname: "Kumar",
			age: 30,
		}
		fmt.Println("Person1:", person1)

		var person2 = new(Person)
		person2.Firstname = "Shyam"
		person2.lastname = "Singh"
		person2.age = 28
		fmt.Println("Person2:", person2)

		var pointer *int
		if(pointer == nil){
			fmt.Println("pointer is nil")
		}
}