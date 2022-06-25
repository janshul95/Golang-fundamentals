package main

import "fmt"

func main() {
	fmt.Println("Methods")
	// as such there is no difference b/w methods and functions
	// normally we call them functions but when they are associated with classes, they are called as methods

	user1 := User{"Anshul", "abv@gmail.com", true, 26}
	fmt.Println("user1", user1)
	fmt.Printf("user's details are %+v \n", user1)
	user1.GetStatus()
	user1.NewMail()
	fmt.Printf("user1: %+v\n", user1)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user loggedin: ", u.Status)
}

// whenever we are passing an object to methods, it is passed as a copy
// that's why we use pointer
func (u User) NewMail()  {
	u.Email = "updated@email.com"
	fmt.Println("email of this user is ", u.Email)
}
