package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("Hello2")
	defer fmt.Println("Hello3")
	fmt.Println("World")
	defer fmt.Println("Hello4")
	myDefer()
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	
}