package main

import (
	"fmt"
)

func main() {
	// in Go there are a few manners of declare variables
	var myVar int // we cen use the keyword 'var', specifying the name of the variable followed by its type
	myVar = 12    // in Go, just like in C, we can declare and initialize a variable in seperated steps

	var anotherVar = 43 // in this case we declared and initialized the variable in a single line and delegated to the compiler the responsibility to handle the type of the variable (automatic typing)

	// oneMorevar int = 64 | on the other hand we can't just declare the variable with its type and hide the 'var' keyword

	oneMoreVar := 64 // finally we have the 'gopher' operator (called the 'short variable declaration' in The Go Programming Language Book) which is a shorthand to declare a variable

	fmt.Println(myVar)
	fmt.Println(anotherVar)
	fmt.Println(oneMoreVar)

	// we can check the type of the variables in this way
	fmt.Printf("myVar: %v | %T\n", myVar, myVar)
	fmt.Printf("anotherVar: %v | %T\n", anotherVar, anotherVar)
	fmt.Printf("oneMoreVar: %v | %T\n", oneMoreVar, oneMoreVar)
}
