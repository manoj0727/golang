package main

import "fmt"

func main(){
	fmt.Println("chal bhai ab array samjha deta hu")

	var arr = []int{10,23,23,43,23};
	fmt.Println(arr);
	var arr1 = [5]string{"manoj"};
	fmt.Println(arr1);
	if arr1[0] == "manoj"{
		fmt.Println("yes name match hua")
	}
	var arr2 = make([]int,5,10);
	arr2[0]=23
	arr2[1]=43
	arr2[2]=53
	arr2 = append(arr2,63,73,83,93,63,73,83,93)
	fmt.Println(arr2)
	fmt.Println(cap(arr2));
}