package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointer")

	// var one int = 2

	// var ptr *int 
	// fmt.Println("value of ptr is ", ptr)

	myNumber := 23
	var ptr *int = &myNumber
	fmt.Println("value of pointer is ",ptr)
	fmt.Println("value ptr is pointing to ",*ptr)

	*ptr = *ptr *2
	fmt.Println("New value is ",myNumber)


}
