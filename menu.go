package main

import(
	"fmt"
)

type Menu struct {
	id int
	name string
	restaurantId string
}

func getMenus(restaurantId int) string {
	return "get menus for " + "yo"
}
func getMenuById(restaurantId int, id int) string {
	return "get menu " + "yo" + " for " + "yo"
}

func postMenu(restaurantId int, data interface{}) string {
	return "add menu for yo"
}

func putMenu(restaurantId int, id int, data interface{}) {
	fmt.Println("update menu id ", id, " for ", restaurantId)
}

func deleteMenu(restaurantId int, id int) {
	fmt.Println("delete menu", id, " for ", restaurantId)
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