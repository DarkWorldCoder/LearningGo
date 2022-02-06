package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to handling urls in golanf")
	// fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	oparams := result.Query()
	fmt.Println("The type : ", oparams)
}
