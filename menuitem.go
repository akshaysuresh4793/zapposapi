package main

import(
	"errors"
	"strconv"
	"strings"
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
	cacheKey := "menuitem_restaurantid_" + strconv.Itoa(restaurantId) + "_menuid_" + strconv.Itoa(menuId) + "_limit_" + strconv.Itoa(limit) + "_offset_" + strconv.Itoa(offset)
	cacheValue := get(cacheKey)
	if len(cacheValue) > 0 {
		return cacheValue
	}
	result, err := readAllMenuItems(restaurantId, menuId, limit, offset)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	set(cacheKey, encode(resp))
	return encode(resp)
}

func getMenuItemById(restaurantId int, menuId int, id int) string {
	var resp Response
	resp.Status = "fail"
	cacheKey := "menuitem_restaurantid_" + strconv.Itoa(restaurantId) + "_menuid_" + strconv.Itoa(menuId) + "_id_" + strconv.Itoa(id)
	cacheValue := get(cacheKey)
	if len(cacheValue) > 0 {
		return cacheValue
	}
	result, err := readMenuItem(restaurantId, menuId, id)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	set(cacheKey, encode(resp))
	return encode(resp)
}

func postMenuItem(restaurantId int, menuId int, m MenuItem) string {
	var resp Response
	var err error
	resp.Status = "fail"
	if restaurantId == 0 {
		err =  errors.New("restaurantId cannot be zero")
		resp.Message = err.Error()
		return encode(resp)
	}
	if menuId == 0 {
		err =  errors.New("menuId cannot be zero")
		resp.Message = err.Error()
		return encode(resp)
	}
	if(len(m.Name) == 0 || len(m.Name) > 500) {
		err = errors.New("name length should be between 1 and 500")
		resp.Message = err.Error()
		return encode(resp)
	}
	if(len(m.Description) == 0 || len(m.Description) > 1000) {
		err = errors.New("description length should be between 1 and 1000")
		resp.Message = err.Error()
		return encode(resp)
	}
	m.RestaurantId = restaurantId
	m.MenuId = menuId
	err = m.create()
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Inserted successfully"
	return encode(resp)
}

func putMenuItem(restaurantId int, menuId int, id int, r map[string]interface{}) string {
	var resp Response
	resp.Status = "fail"
	err := updateMenuItem(id, r)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Updated successfully"
	cacheKey := "menuitem_restaurantid_" + strconv.Itoa(restaurantId) + "_menuid_" + strconv.Itoa(menuId) + "_id_" + strconv.Itoa(id)
	delete(cacheKey)
	return encode(resp)
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
	cacheKey := "menuitem_restaurantid_" + strconv.Itoa(restaurantId) + "_menuid_" + strconv.Itoa(menuId) + "_id_" + strconv.Itoa(id)
	delete(cacheKey)
	return encode(resp)
}

// create a menu-item
func(m *MenuItem) create() error {
	result, err := db.Exec("INSERT INTO menu_item(name, description, restaurant_id, menu_id) VALUES (?, ?, ?, ?)", m.Name, m.Description, m.RestaurantId, m.MenuId)
	handleError(err)
	if err != nil {
		return err
	}
	_ = result
	return err
}

func readMenuItem(restaurantId int, menuId int, id int) (MenuItem, error) {
	var result MenuItem
	var Name string
	var Description string
	var Id int
	prepare, err := db.Prepare("SELECT id, name, description FROM menu_item WHERE id = ? AND restaurant_id = ? AND menu_id = ?")
	handleError(err)
	if err != nil {
		return result, err
	}
	err = prepare.QueryRow(id, restaurantId, menuId).Scan(&Id, &Name, &Description)
	handleError(err)
	if err != nil {
		return result, err
	}
	result.Id = Id
	result.Name = Name
	result.RestaurantId = restaurantId
	result.MenuId = menuId
	result.Description = Description
	return result, err
}

func readAllMenuItems(restaurantId int, menuId int, limit int, offset int) ([]MenuItem, error) {
	var result []MenuItem
	var m MenuItem
	var Name string
	var Description string
	var Id int
	prepare, err := db.Query("SELECT id, name, description FROM menu_item WHERE restaurant_id = ? AND menu_id = ? LIMIT ?,?", restaurantId, menuId, limit, offset)
	handleError(err)
	if err != nil {
		return result, err
	}
	for prepare.Next() {
		err = prepare.Scan(&Id, &Name, &Description)
		handleError(err)
		if err != nil {
			return result, err
		}
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
func updateMenuItem(id int, r map[string] interface{}) error {
	var err error
	var restaurantId int
	var menuId int
	output := "UPDATE menu_item SET "
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
		if k == "description" { 
			if v.(string) == "0" {
				err = errors.New("Description has to be between 1 and 500 characters")
				return err
			} else {
				output += "description = ?, "
				values = append(values, v.(string))
			}
		}
		if k == "menuId" { 
			if v.(string) == "0" {
				err = errors.New("menuId cannot be 0")
				return err
			} else {
				output += "menu_id = ?, "
				menuId, err = strconv.Atoi(v.(string))
				handleError(err)
				if err != nil {
					return err
				}
				values = append(values, menuId)
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