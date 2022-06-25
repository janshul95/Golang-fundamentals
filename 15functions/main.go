package main

import "fmt"

// main acts as the entrypoint in the program
func main() {
	fmt.Println("Functions in golang")
	greeter()
	greeter2()

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proresult, msg := proAdder(23, 5354, 6565, 768, 7, 3, 5)
	fmt.Println("Result is ", proresult, msg)
}

func adder(i1 int, i2 int) int {
	return i1 + i2
}

// passing multiple no of values
func proAdder(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}

	return total, "test string"

}

// we are not allowed to write function inside a function
func greeter() {
	fmt.Println("Namaste Duniya")
}
func greeter2() {
	fmt.Println("Another function")
}
