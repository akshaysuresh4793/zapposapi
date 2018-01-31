package main

import(
	"fmt"
)

type MenuItem struct {
	id int
	name string
	menu_id string
	description string
}

// create a menu-item
func(m *MenuItem) create() {
	fmt.Println("Create");
}

// read a menu-item
func (m *MenuItem) read(){
	fmt.Println("Read");
}

// update a menu-item
func (m *MenuItem) update(){
	fmt.Println("Update");
}

// delete a menu-item
func (m *MenuItem) delete() {
	fmt.Println("Delete");
}