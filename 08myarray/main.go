package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go")
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("fruit list is",fruitList)
	fmt.Println("length is", len(fruitList))

	var vegList =[3]string{"potato","beans","mushroom"}
	fmt.Println("veg list is ",vegList)
}
