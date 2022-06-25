package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	msg := "Please give an input"
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for pizza:")

	// comma ok || err of syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T",input)

}
