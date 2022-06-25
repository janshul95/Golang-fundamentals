package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all languages: ", languages)
	fmt.Println("JS shorts for ", languages["JS"])

	// delete a value
	delete(languages, "RB")
	fmt.Println("list of all languages: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf(" value is %v\n", value)
	}
}
