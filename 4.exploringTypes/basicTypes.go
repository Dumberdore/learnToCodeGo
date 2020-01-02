package main

import "fmt"

var x = 42
var y string = "Shaken, not Stirred!"

// Declare VARIABLE of certain type, in this case String and Assign value
var z string = `James Said, Shaken..

not stirred!!`

// This is STATIC Programming language
// A VARIABLE is DECLARED to hold a VALUE of certain TYPE

func main() {
	fmt.Println("x : ", x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
