package main

import(
	"fmt"
	"encoding/json"
)

type Person struct{
	Name  string `json:"full_name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main(){
	fmt.Println("karo")

	person := Person{
		Name:  "Manoj Kumawat",
		Age:   24,
		Email: "manoj@gmail.com",
	}
	
	fmt.Printf("Person: %+v\n", person)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	
	fmt.Println("JSON Data:", string(jsonData))
}