package main
import(
"fmt"
"time"
) 

func main(){
	defer fmt.Println("manoj ")
	currentTime := time.Now();
	fmt.Println("Current time is:", currentTime)
	fomattedTime := currentTime.Format("02-01-2006 15:04 PM") // this is fix use this format that can set format this is go official release time of go lang
	fmt.Println("Formatted time is:", fomattedTime)
}