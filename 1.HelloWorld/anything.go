package main

import "fmt"

func main() {
	//contorl flow
	// 1: seq
	// 2. loop
	// 3. iterative

	fmt.Println("Hello world at Go!")
	foo()

	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in Foo!")
}

func bar() {
	fmt.Println("Control Flow Babeyyy!!!")
}
