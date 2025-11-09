package main

import "fmt"

func main(){
	fmt.Println("chal bhai ab array samjha deta hu")

	var arr = [5]int{10,23,23,43,23};
	fmt.Println(arr);
	var arr1 = [5]string{"manoj"};
	fmt.Println(arr1);
	if arr1[0] == "manoj"{
		fmt.Println("yes name match hua")
	}
}