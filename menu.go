package main

import(
	"fmt"
)

type Menu struct {
	id int
	name string
	restaurant_id string
}

// create a menu
func(m *Menu) create() {
	fmt.Println("Create");
}

// read a menu
func (m *Menu) read(){
	fmt.Println("Read");
}

// update a menu
func (m *Menu) update(){
	fmt.Println("Update");
}

// delete a menu
func (m *Menu) delete() {
	fmt.Println("Delete");
}