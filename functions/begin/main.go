// functions/begin/main.go
package main

import "fmt"

// simple greet function
func greet() string {
	return "Hello!"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	return "Hello " + name + "!"
}

// greetWithName returns a greeting with the name and age of the person
func greetWithNameAndAge(name string, age int) string {
	return fmt.Sprintf("Yo %s (%d)!", name, age)
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("divide by zero")
	}

	return n1/n2, nil
} 

func main() {
	// invoke greet function
	fmt.Println(greet())

	// invoke greetWithName function
	fmt.Println(greetWithName("Toni"))
	fmt.Println(greetWithNameAndAge("Bessie", 31))

	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
