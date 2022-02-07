package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "https://randomuser.me/api/"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	data := string(content)
	fmt.Println(data)
	// fmt.Println(data.results)
}
