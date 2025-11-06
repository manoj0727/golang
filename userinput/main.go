package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	fmt.Println("hey how are you?")
	// var name string;

	// fmt.Scan(&name)
	// fmt.Println("i am also",name);

	// get input of complete line thats it
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("hello, Mr", name)

}