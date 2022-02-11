package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	PerformPostJson()
}

func PerformPostJson() {
	const myurl = ""

	requestBody := strings.NewReader(`
	{
		"courseName":"Lets go",
		"price":0.
		"platform":"online"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormReq() {
	const myurl = ""

	data := url.Values{}

	data.Add("firstname", "ayush")
	data.Add("lastname", "niroula")
	data.Add("email", "ayush@gmail.com")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
