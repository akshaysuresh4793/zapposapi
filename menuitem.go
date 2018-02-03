package main

import(
	"fmt"
)

type MenuItem struct {
	id int
	name string
	menuId int
	restaurantId int
	description string
}

func getMenuItems(restaurantId int, menuId int) string {
	return "menu-items for " + "yo" + " from " + "yo"
}

func getMenuItemById(restaurantId int, menuId int, id int) string {
	return "menu-item id " + "yo" + " for " + "yo" + " from " + "yo"
}

func postMenuItem(restaurantId int, menuId int, data interface{}) string {
	return "add menu item for yo from yo"
}

func putMenuItem(restaurantId int, menuId int, id int, data interface{}) {
	fmt.Println("update menu-item ", id, " of menu ", menuId, " and restaurantId ", restaurantId)
}

func deleteMenuItem(restaurantId int, menuId int, id int) {
	fmt.Println("delete menu item ", id, " for ", menuId, " from ", restaurantId)
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