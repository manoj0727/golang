package main

import(
	"fmt"
	"os"
	"io"
)

func main(){
	file,err := os.Create("testfile.txt")
	if err != nil{
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	content := "Hello, this is a test file.\nWelcome to Golang file handling."

	byte, errors := io.WriteString(file,content)
	fmt.Printf("Wrote %d bytes to file.\n", byte)

	if errors != nil{
		fmt.Println("Error writing to file:", errors)
		return
	}
}