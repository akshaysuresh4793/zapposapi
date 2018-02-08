package main

import (
	"errors"
	_ "fmt"
	"strconv"
	"strings"
)

type Restaurant struct {
	Id         int    `json:"id,string"`
	Name       string `json:"name"`
	LocationId int    `json:"locationId,string"`
}

func getRestaurants(limit int, offset int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readAllRestaurants(limit, offset)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	return encode(resp)
}

func getRestaurantById(id int) string {
	var resp Response
	resp.Status = "fail"
	cacheKey := "restaurant_id_" + strconv.Itoa(id)
	cacheValue := get(cacheKey)
	if len(cacheValue) > 0 {
		return cacheValue
	}
	result, err := readRestaurant(id)
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

func postRestaurant(r Restaurant) string {
	var resp Response
	resp.Status = "fail"
	var err error
	if r.Id != 0 {
		// id cannot be non-zero this is a new insert
		err = errors.New("Id is nonzero")
		resp.Message = err.Error()
		return encode(resp)
	}
	if len(r.Name) == 0 || len(r.Name) > 500 {
		// MySQL has varchar(500) - this would break
		err = errors.New("Name should be between 1 and 500 characters")
		resp.Message = err.Error()
		return encode(resp)
	}
	if r.LocationId == 0 {
		// location cannot be empty
		err = errors.New("Location cannot be empty")
		resp.Message = err.Error()
		return encode(resp)
	}

	// all okay
	err = r.create()
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Inserted successfully"
	return encode(resp)
}

func putRestaurant(id int, r map[string]interface{}) string {
	var resp Response
	resp.Status = "fail"
	err := updateRestaurant(id, r)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Updated successfully"
	cacheKey := "restaurant_id_" + strconv.Itoa(id)
	delete(cacheKey)
	return encode(resp)
}

func deleteRestaurant(id int) string {
	var resp Response
	resp.Status = "fail"
	var r Restaurant
	r.Id = id
	err := r.delete()
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Deleted successfully"
	cacheKey := "restaurant_id_" + strconv.Itoa(id)
	delete(cacheKey)
	return encode(resp)
}

// create a restaurant
func (r *Restaurant) create() error {
	// prepare MySQL query
	// insert
	result, err := db.Exec("INSERT INTO restaurant(name, location_id) VALUES (?, ?)", r.Name, r.LocationId)
	handleError(err)
	_ = result
	return err
}

// read a restaurant
func readRestaurant(id int) (Restaurant, error) {
	var result Restaurant
	var Name string
	var Id int
	var Location int
	prepare, err := db.Prepare("SELECT id, name, location_id FROM restaurant WHERE id = ?")
	handleError(err)
	if err != nil {
		return result, err
	}
	err = prepare.QueryRow(id).Scan(&Id, &Name, &Location)
	handleError(err)
	if err != nil {
		return result, err
	}
	result.Id = Id
	result.Name = Name
	result.LocationId = Location
	return result, err
}

func readAllRestaurants(limit int, offset int) ([]Restaurant, error) {
	var result []Restaurant
	var r Restaurant
	var Name string
	var Id int
	var Location int
	prepare, err := db.Query("SELECT id, name, location_id FROM restaurant LIMIT ?,?", limit, offset)
	handleError(err)
	if err != nil {
		return result, err
	}
	for prepare.Next() {
		err = prepare.Scan(&Id, &Name, &Location)
		handleError(err)
		if err != nil {
			return result, err
		}
		r.Id = Id
		r.Name = Name
		r.LocationId = Location
		result = append(result, r)
	}
	return result, err
}

func updateRestaurant(id int, r map[string]interface{}) error {
	var err error
	var locationId int
	output := "UPDATE restaurant SET "
	var values []interface{}
	for k, v := range r {
		if k == "name" {
			if len(v.(string)) == 0 || len(v.(string)) > 500 {
				err = errors.New("Name has to be between 1 and 500 characters")
				return err
			} else {
				output += "name = ?, "
				values = append(values, v.(string))
			}
		}
		if k == "locationId" {
			if v.(string) == "0" {
				err = errors.New("locationId cannot be 0")
				return err
			} else {
				output += "location_id = ?, "
				locationId, err = strconv.Atoi(v.(string))
				handleError(err)
				if err != nil {
					return err
				}
				values = append(values, locationId)
			}
		}
	}
	output = strings.Trim(output, ", ")
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

// delete a restaurant
func (r *Restaurant) delete() error {
	prepare, err := db.Exec("DELETE FROM menu_item WHERE restaurant_id = ?", r.Id)
	handleError(err)
	if err != nil {
		return err
	}
	prepare, err = db.Exec("DELETE FROM menu WHERE restaurant_id = ?", r.Id)
	handleError(err)
	if err != nil {
		return err
	}
	prepare, err = db.Exec("DELETE FROM restaurant WHERE id = ?", r.Id)
	handleError(err)
	if err != nil {
		return err
	}
	_ = prepare
	// update the cache
	return err
}
