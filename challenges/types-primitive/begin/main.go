// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 100

	// print the value of the local variable "val"
	fmt.Println(val)

	// print the value of the package-level variable "val"
	fmt.Println(getPackageLevelVal())

	// update the package-level variable "val" and print it again
	updatePackageLevelVal()
	fmt.Println(getPackageLevelVal())

	// print the pointer address of the local variable "val"
	fmt.Println(&val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 123
	fmt.Println(val)
}

func getPackageLevelVal() string {
	return val
}

func updatePackageLevelVal() {
	val = "not so global"
}