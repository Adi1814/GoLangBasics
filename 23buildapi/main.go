package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for courses

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db

var courses []Course

//middlewares and helper fucntions

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("APIs in golang!")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "cats", CoursePrice: 6734, Author: &Author{Fullname: "Aditi Mishra", Website: "jhbfhwsgf@meow.com"}})

	//listen to a port
	log.Fatal(http.ListenAndServe(":7474", r))

}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by aditi and bruno</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("conetent-type", "application/json")
	json.NewEncoder(w).Encode(courses) //seeding when we want to test our db we put in dump values in that

}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course ")
	w.Header().Set("conetent-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Get one course ")
	w.Header().Set("conetent-type", "application/json")
	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what id data is being sent like this "{}"
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	//generate a unique id
	//convert that id into string
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course ")
	w.Header().Set("conetent-type", "application/json")

	//lets grab id from request
	params := mux.Vars(r)

	//loop through id remove, add in id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.Encoder(w).Encode(course)
			return
		}
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if params["id"] == "" {
		json.NewEncoder(w).Encode("id not found")
		return
	}
	//send a response when id is not found
}
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course ")
	w.Header().Set("conetent-type", "application/json")

	//lets grab id from request
	params := mux.Vars(r)

	//loop, id, remove(index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
