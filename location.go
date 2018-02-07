package main

import(
	"errors"
	"strings"
	"strconv"
	"fmt"
)

type Location struct {
	Id int `json: "id,string"`
	Name string `json:"name"`
}

func getLocations(limit int, offset int) string {
	var resp Response
	resp.Status = "fail"
	cacheKey := "location_" + strconv.Itoa(limit) + "_offset_" + strconv.Itoa(offset)
	cacheValue := get(cacheKey)
	if len(cacheValue) > 0 {
		return cacheValue
	}
	result, err := readAllLocations(limit, offset)
	if err != nil {
		resp.Message = err.Error()
		return  encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	set(cacheKey, encode(resp))
	return encode(resp)
}

func getLocationById(id int) string {
	var resp Response
	resp.Status = "fail"
	cacheKey := "location_id_" + strconv.Itoa(id)
	cacheValue := get(cacheKey)
	if len(cacheValue) > 0 {
		return cacheValue
	}
	result, err := readLocation(id)
	if (err != nil) {
		resp.Message = err.Error()
		return  encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	set(cacheKey, encode(resp))
	return encode(resp)
}

func postLocation(l Location) string {
	var resp Response
	resp.Status = "fail"
	// data validation
	if(l.Id != 0) {
		err := errors.New("Id is nonzero")
		// id cannot be non-zero this is a new insert
		resp.Message = err.Error()
		return  encode(resp)
	}
	if len(l.Name) == 0 && len(l.Name) > 500 {
		err := errors.New("Name should be between 1 and 500 characters")
		// MySQL has varchar(500) - this would break
		resp.Message = err.Error()
		return  encode(resp)
	}
	// TODO: check if the name already exists
	err := l.create()
	if err != nil {
		resp.Message = err.Error()
		return  encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Inserted successfully"
	return encode(resp)
}

func putLocation(id int, r map[string]interface{}) string {
	var resp Response
	resp.Status = "fail"
	err := updateLocation(id, r)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Updated successfully"
	cacheKey := "location_id_" + strconv.Itoa(id)
	delete(cacheKey)
	return encode(resp)
}

func deleteLocation(id int) string {
	var resp Response
	resp.Status = "fail"
	var l Location
	l.Id = id
	err := l.delete()
	if err != nil {
		resp.Message = err.Error()
		return  encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Deleted successfully"
	cacheKey := "location_id_" + strconv.Itoa(id)
	delete(cacheKey)
	return encode(resp)
}

// create a location
func(l *Location) create() error {
	result, err := db.Exec("INSERT INTO location(name) VALUES (?)", l.Name)
	handleError(err)
	_ = result
	return err
}

// read a location
func readLocation(id int) (Location, error) {
	var result Location
	var Name string
	var Id int
	prepare, err := db.Prepare("SELECT id, name FROM location WHERE id = ?")
	handleError(err)
	if err != nil {
		return result, err
	}
	err = prepare.QueryRow(id).Scan(&Id, &Name)
	handleError(err)
	result.Id = Id
	result.Name = Name
	return result, err
}


func readAllLocations(limit int, offset int) ([]Location, error) {
	var result []Location
	var l Location
	var Name string
	var Id int
	prepare, err := db.Query("SELECT id, name FROM location LIMIT ?,?", limit, offset)
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
		l.Id = Id
		l.Name = Name
		result = append(result, l)
	}
	return result, err
}

// update a location
func updateLocation(id int, r map[string] interface{}) error {
	var err error
	output := "UPDATE location SET "
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

// delete a location
func (l *Location) delete() error {
	prepare, err := db.Exec("DELETE FROM menu_item WHERE restaurant_id IN (SELECT r.id FROM restaurant r WHERE r.location_id = ?)", l.Id)
	handleError(err)
	if err != nil {
		return err
	}
	prepare, err = db.Exec("DELETE FROM menu WHERE restaurant_id IN (SELECT r.id FROM restaurant r WHERE r.location_id = ?)", l.Id)
	handleError(err)
	if err != nil {
		return err
	}
	prepare, err = db.Exec("DELETE FROM restaurant WHERE location_id = ?", l.Id)
	handleError(err)
	if err != nil {
		return err
	}
	prepare, err = db.Exec("DELETE FROM location WHERE id = ?", l.Id)
	handleError(err)
	if err != nil {
		return err
	}
	_ = prepare
	return err
}