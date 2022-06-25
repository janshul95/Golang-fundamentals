package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")
	rand.Seed(time.Now().UnixNano())
	diceNmber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNmber)

	switch diceNmber {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough //also move on to next case
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6")
	default:
		fmt.Println("What was that")

	}
}
