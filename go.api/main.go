package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

var courses []Course

//middleware

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {

}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// what if: body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Plese send some data")

	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return

	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
