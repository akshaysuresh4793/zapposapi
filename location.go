package main

import(
	"fmt"
)

type Location struct {
	id int
	name string
}

func getLocations() string {
	return "get all locations"
}

func getLocationById(id int) string {
	return "get location number"
}

func postLocation(data interface{}) {
	fmt.Println("add location")
}

func putLocation(id int, data interface{}) {
	fmt.Println("update locations")
}

func deleteLocation(id int) {
	fmt.Println("delete location")
}

// create a location
func(l *Location) create() {
	fmt.Println("Create");
}

// read a location
func (l *Location) read(){
	fmt.Println("Read");
}

// update a location
func (l *Location) update(){
	fmt.Println("Update");
}

// delete a location
func (l *Location) delete() {
	fmt.Println("Delete");
}