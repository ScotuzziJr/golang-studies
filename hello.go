package main // in go, all codes are divided in packages; the main package is special because tells the compiler this is a 
			// "standalone executable program, not a library" -- quote from "The Go Programming Language by Donovan and Kernighan

import "fmt" // this is a external package that handles formated output and input

// entry point of the programa
func main() {
	fmt.Println("hello, world!") // the Println function print stuff on the console with \n in the end
}
