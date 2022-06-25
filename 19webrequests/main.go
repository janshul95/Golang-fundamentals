package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("Web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	databytes,err :=ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println("",content)
	defer response.Body.Close() //caller's responsibility to close the collection

}
