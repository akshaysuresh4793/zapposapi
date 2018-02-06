package main

import(
	"fmt"
	"errors"
)

type Menu struct {
	Id int `json:"id,string"`
	Name string `json:"name"`
	RestaurantId int `json:"restaurantId,string"`
}

func getMenus(restaurantId int, limit int, offset int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readAllMenus(restaurantId, limit, offset)
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
func getMenuById(restaurantId int, id int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readMenu(restaurantId, id)
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
func postMenu(restaurantId int, m Menu) string {
	var resp Response
	var err error
	resp.Status = "fail"
	if(m.Id != 0) {
		// id cannot be non-zero this is new data
		err = errors.New("id cannot be nonzero")
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

	if(restaurantId == 0) {
		err = errors.New("restaurantId cannot be zero")
		handleError(err)
		resp.Message = err.Error()
		return encode(resp)
	}

	m.RestaurantId = restaurantId
	// TODO: check if the restaurant exists

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

func putMenu(restaurantId int, id int, data interface{}) {
	fmt.Println("update menu id ", id, " for ", restaurantId)
}

func deleteMenu(restaurantId int, id int) string {
	var resp Response
	resp.Status = "fail"
	var m Menu
	m.Id = id
	m.RestaurantId = restaurantId
	err := m.delete()
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Deleted successfully"
	return encode(resp)
}

// create a menu
func(m *Menu) create() error {
	// prepare MySQL query
	// insert
	result, err := db.Exec("INSERT INTO menu(name, restaurant_id) VALUES (?, ?)", m.Name, m.RestaurantId)
	handleError(err)
	_ = result
	// update cache
	return err
}


func readMenu(restaurantId int, id int) (Menu, error) {
	prepare, err := db.Prepare("SELECT id, name FROM menu WHERE id = ? AND restaurant_id = ?")
	var result Menu
	var Name string
	var Id int
	err = prepare.QueryRow(id, restaurantId).Scan(&Id, &Name)
	handleError(err)
	result.Id = Id
	result.Name = Name
	result.RestaurantId = restaurantId
	return result, err
}


func readAllMenus(restaurantId int, limit int, offset int) ([]Menu, error) {
	prepare, err := db.Query("SELECT id, name FROM menu WHERE restaurant_id = ? LIMIT ?,?", restaurantId, limit, offset)
	var result []Menu
	var m Menu
	var Name string
	var Id int
	for prepare.Next() {
		err = prepare.Scan(&Id, &Name)
		handleError(err)
		m.Id = Id
		m.Name = Name
		m.RestaurantId = restaurantId
		result = append(result, m)
	}
	return result, err
}

// update a menu
func (m *Menu) update(){
	fmt.Println("Update");
}

// delete a menu
func (m *Menu) delete() error {
	prepare, err := db.Exec("DELETE FROM menu_item WHERE menu_id = ? AND restaurant_id = ?", m.Id, m.RestaurantId)
	handleError(err)
	if err != nil {
		return err
	}
	prepare, err = db.Exec("DELETE FROM menu WHERE id = ? AND restaurant_id = ?", m.Id, m.RestaurantId)
	handleError(err)
	if err != nil {
		return err
	}
	_ = prepare
	// update cache

	return err
}