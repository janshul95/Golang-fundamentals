package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Implementation for http verbs ")
	PerformGetRequest()
	PerformPOSTRequest()
	PerformFORMRequest()
}

func PerformFORMRequest() {
	fmt.Println("Initiating POST request...")
	defer fmt.Println("-----Competed the request")
	const myURL = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "anshul")
	data.Add("latname", "jain")

	response, err := http.PostForm(myURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("", string(content))

}
func PerformPOSTRequest() {
	fmt.Println("Initiating POST request")
	defer fmt.Println("Competed the request")
	const myURL = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"name":"anshul",
		"project":"go with go"
	}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("", string(content))
}

func PerformGetRequest() {
	fmt.Println("Initiating get request ...")
	defer fmt.Println("Get request completed ...")
	const getURL = "http://localhost:8000/get"

	response, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)
	fmt.Println("byte count", byteCount)

	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}
