package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in golang")

	content := "this needs to go into a file ...."

	file, err := os.Create("./test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is ", length)
	file.Close()

	readFile("./test.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is\n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}
