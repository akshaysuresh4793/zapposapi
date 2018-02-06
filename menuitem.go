package main

import(
	"fmt"
	"errors"
)

type MenuItem struct {
	Id int `json:"id,string"`
	Name string `json:"name"`
	MenuId int `json:"menuId,string"`
	RestaurantId int `json:"restaurantId,string"`
	Description string `json:"description"`
}

func getMenuItems(restaurantId int, menuId int, limit int, offset int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readAllMenuItems(restaurantId, menuId, limit, offset)
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	return encode(resp)
}

func getMenuItemById(restaurantId int, menuId int, id int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readMenuItem(restaurantId, menuId, id)
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	return encode(resp)
}

func postMenuItem(restaurantId int, menuId int, m MenuItem) string {
	var resp Response
	var err error
	resp.Status = "fail"
	if restaurantId == 0 {
		err =  errors.New("restaurantId cannot be zero")
		handleError(err)
		resp.Message = err.Error()
		return encode(resp)
	}
	if menuId == 0 {
		err =  errors.New("menuId cannot be zero")
		handleError(err)
		resp.Message = err.Error()
		return encode(resp)
	}
	if(len(m.Name) == 0 || len(m.Name) > 500) {
		err = errors.New("name length should be between 1 and 500")
		handleError(err)
		resp.Message = err.Error()
		return encode(resp)
	}
	if(len(m.Description) == 0 || len(m.Description) > 1000) {
		err = errors.New("description length should be between 1 and 1000")
		handleError(err)
		resp.Message = err.Error()
		return encode(resp)
	}
	// TODO: check if the restaurant exists
	m.RestaurantId = restaurantId

	// TODO: check if the menu exists
	m.MenuId = menuId
	err = m.create()
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Inserted successfully"
	return encode(resp)
}

func putMenuItem(restaurantId int, menuId int, id int, data interface{}) {
	fmt.Println("update menu-item ", id, " of menu ", menuId, " and restaurantId ", restaurantId)
}

func deleteMenuItem(restaurantId int, menuId int, id int) string {
	var m MenuItem
	var resp Response
	resp.Status = "fail"
	m.Id = id
	m.MenuId = menuId
	m.RestaurantId = restaurantId
	err := m.delete()
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Deleted successfully"
	return encode(resp)
}

// create a menu-item
func(m *MenuItem) create() error {
	result, err := db.Exec("INSERT INTO menu_item(name, description, restaurant_id, menu_id) VALUES (?, ?, ?, ?)", m.Name, m.Description, m.RestaurantId, m.MenuId)
	if err != nil {
		return err
	}
	_ = result
	// update cache
	return nil
}

func readMenuItem(restaurantId int, menuId int, id int) (MenuItem, error) {
	prepare, err := db.Prepare("SELECT id, name, description FROM menu_item WHERE id = ? AND restaurant_id = ? AND menu_id = ?")
	var result MenuItem
	var Name string
	var Description string
	var Id int
	err = prepare.QueryRow(id, restaurantId, menuId).Scan(&Id, &Name, &Description)
	handleError(err)
	result.Id = Id
	result.Name = Name
	result.RestaurantId = restaurantId
	result.MenuId = menuId
	result.Description = Description
	return result, err
}

func readAllMenuItems(restaurantId int, menuId int, limit int, offset int) ([]MenuItem, error) {
	prepare, err := db.Query("SELECT id, name, description FROM menu_item WHERE restaurant_id = ? AND menu_id = ? LIMIT ?,?", restaurantId, menuId, limit, offset)
	var result []MenuItem
	var m MenuItem
	var Name string
	var Description string
	var Id int
	for prepare.Next() {
		err = prepare.Scan(&Id, &Name, &Description)
		handleError(err)
		m.Id = Id
		m.Name = Name
		m.RestaurantId = restaurantId
		m.MenuId = menuId
		m.Description = Description
		result = append(result, m)
	}
	return result, err
}

// update a menu-item
func (m *MenuItem) update(){
	fmt.Println("Update");
}

// delete a menu-item
func(m *MenuItem) delete() error {
	prepare, err := db.Exec("DELETE FROM menu_item WHERE id = ? AND restaurant_id = ? AND menu_id = ?", m.Id, m.RestaurantId, m.MenuId)
	handleError(err)
	if err != nil {
		return err
	}
	_ = prepare
	// update cache
	return err
}