// packages/basics/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\nYello, world!\n")
	fmt.Printf("It's such a lovely %s!\n", time.Now().Weekday())
}
