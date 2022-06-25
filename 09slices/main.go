package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("fruit list is", fruitList)
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println("", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("", fruitList)

	highScores := make([]int,4)

	highScores[0] = 2232
	highScores[1] = 121
	highScores[2]=33243
	highScores[3]=43243

	highScores = append(highScores, 34234,555)

	fmt.Println("",highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// remove a value from slices

	var courses = []string{"reactjs","javascipt","not-required","swift",}
	fmt.Println("",courses)

	index :=2

	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("",courses)


}
