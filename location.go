package main

import(
	"fmt"
)

type Location struct {
	id int
	name string
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