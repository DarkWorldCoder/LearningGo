package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	price    int
	platform string
	password string
	tags     []string
}

func main() {
	EncodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"react js", 200, "learn", "abc21", []string{"web-dev", "js"}},
		{"mern js", 100, "learn", "abc21", []string{"fullstack", "js"}},
		{"angular js", 400, "learn", "abc21", []string{"web-dev", "js"}},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	// json.Unmarshal()
}
