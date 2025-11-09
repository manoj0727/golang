package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	
)
func main(){
	fmt.Println("web request in golang")
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if(err != nil){
		fmt.Println("error occured:",err)
		return
	}
	defer res.Body.Close()
	fmt.Println("response status:",res.Status)

	data,err := ioutil.ReadAll(res.Body)
	if(err != nil){
		fmt.Println("error reading response body:",err)
		return
	}
	fmt.Println("response body:",string(data))
}