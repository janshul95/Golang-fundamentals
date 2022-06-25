package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://bing.com:3000/search?q1=dotnet&q2=dotnet&q3=dotnet"

func main() {
	fmt.Println("handling urls in go")
	fmt.Println(myurl)

	u, _ := url.Parse(myurl)
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Query())
	fmt.Println(u.Port())

	qparams := u.Query()
	fmt.Printf("the types of query params are %T\n",qparams)
	fmt.Println(qparams["q1"])

	for _, v := range qparams {
		fmt.Println("Param is : ",v)
	}

	partsOfURL := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=anshul",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
