package helper 

import "fmt"

//When you capitalize the name of the function, it gets exported
func JustPrint() {
	fmt.Println("Just printing from helper package's helper.go")
}