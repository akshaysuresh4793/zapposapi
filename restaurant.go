package main

import(
	"fmt"
)

type Restaurant struct {
	id int
	name string
	location string
}

// create a restaurant
func(r *Restaurant) create() {
	fmt.Println("Create");
}

// read a restaurant
func (r *Restaurant) read(){
	fmt.Println("Read");
}

// update a restaurant
func (r *Restaurant) update(){
	fmt.Println("Update");
}

// delete a restaurant
func (r *Restaurant) delete() {
	fmt.Println("Delete");
}