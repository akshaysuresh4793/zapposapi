package main

import(
	"errors"
	"strconv"
	"strings"
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
		resp.Message = err.Error()
		return encode(resp)
	}

	if(len(m.Name) == 0 || len(m.Name) > 500) {
		err = errors.New("name length should be between 1 and 500")
		resp.Message = err.Error()
		return encode(resp)
	}

	if(restaurantId == 0) {
		err = errors.New("restaurantId cannot be zero")
		resp.Message = err.Error()
		return encode(resp)
	}

	m.RestaurantId = restaurantId
	// TODO: check if the restaurant exists

	err = m.create()
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Inserted successfully"
	return encode(resp)
}

func putMenu(restaurantId int, id int, r map[string]interface{}) string {
	var resp Response
	resp.Status = "fail"
	err := updateMenu(id, r)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Updated successfully"
	return encode(resp)
}

func deleteMenu(restaurantId int, id int) string {
	var resp Response
	resp.Status = "fail"
	var m Menu
	m.Id = id
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
	var result Menu
	var Name string
	var Id int
	prepare, err := db.Prepare("SELECT id, name FROM menu WHERE id = ? AND restaurant_id = ?")
	handleError(err)
	if err != nil {
		return result, err
	}
	err = prepare.QueryRow(id, restaurantId).Scan(&Id, &Name)
	handleError(err)
	if err != nil {
		return result, err
	}
	result.Id = Id
	result.Name = Name
	result.RestaurantId = restaurantId
	return result, err
}


func readAllMenus(restaurantId int, limit int, offset int) ([]Menu, error) {
	var result []Menu
	var m Menu
	var Name string
	var Id int
	prepare, err := db.Query("SELECT id, name FROM menu WHERE restaurant_id = ? LIMIT ?,?", restaurantId, limit, offset)
	handleError(err)
	if err != nil {
		return result, err
	}
	for prepare.Next() {
		err = prepare.Scan(&Id, &Name)
		handleError(err)
		if err != nil {
			return result, err
		}
		m.Id = Id
		m.Name = Name
		m.RestaurantId = restaurantId
		result = append(result, m)
	}
	return result, err
}

// update a menu
func updateMenu(id int, r map[string] interface{}) error {
	var err error
	var restaurantId int
	output := "UPDATE menu SET "
	var values []interface{}
	for k,v := range(r) {
		if k == "name" {
			if len(v.(string)) == 0 || len(v.(string)) > 500 {
				err = errors.New("Name has to be between 1 and 500 characters")
				return err
			} else {
				output += "name = ?, "
				values = append(values, v.(string))
			}
		}
		if k == "restaurantId" { 
			if v.(string) == "0" {
				err = errors.New("restaurantId cannot be 0")
				return err
			} else {
				output += "restaurant_id = ?, "
				restaurantId, err = strconv.Atoi(v.(string))
				handleError(err)
				if err != nil {
					return err
				}
				values = append(values, restaurantId)
			}
		}
	}
	output = strings.Trim(output,", ")
	output += " WHERE id = ?"
	values = append(values, id)
	prepare, errr := db.Exec(output, values...)
	handleError(errr)
	_ = prepare
	if errr != nil {
		return errr
	}
	return err
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