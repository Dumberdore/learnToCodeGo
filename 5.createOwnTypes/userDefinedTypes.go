package main

import "fmt"

var x int

type hotdog int

var b hotdog

func main() {
	x = 10
	b = 25

	fmt.Printf("%d\t%T\n", x, x)
	fmt.Printf("%d\t%T\n", b, b)

	x = int(b)
	fmt.Printf("%d\t%T\n", x, x)
}
