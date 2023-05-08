// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	var myFloat = 17.05
	var myInt = uint(myFloat)

	fmt.Println("myFloat, myInt:", myFloat, myInt)
}
