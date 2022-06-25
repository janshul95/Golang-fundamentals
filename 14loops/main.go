package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println("days", days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])

	}

	for i, v := range days {
		fmt.Printf("index is %v and value is %v\n",i,v)
	}

	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 5 {
			break
		}
		
		fmt.Println("value is ", rougeValue)
		rougeValue++
	}
}
