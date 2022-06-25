package main

import "fmt"

func main() {
	fmt.Println("Struct in Golang")
	// no inheritance in golang; No super or parent

	user1 := User{"Anshul", "abv@gmail.com", true, 26}
	fmt.Println("user1", user1)
	fmt.Printf("user's details are %+v \n", user1)
	fmt.Printf("user's name is %v and email is %v", user1.Name, user1.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
