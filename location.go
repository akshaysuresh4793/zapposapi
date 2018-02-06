package main

import(
	"fmt"
	"errors"
)

type Location struct {
	Id int `json: "id,string"`
	Name string `json:"name"`
}

func getLocations(limit int, offset int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readAllLocations(limit, offset)
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return  encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	return encode(resp)
}

func getLocationById(id int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readLocation(id)
	handleError(err)
	if (err != nil) {
		resp.Message = err.Error()
		return  encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	return encode(resp)
}

func postLocation(l Location) string {
	var resp Response
	resp.Status = "fail"
	// data validation
	if(l.Id != 0) {
		err := errors.New("Id is nonzero")
		handleError(err)
		// id cannot be non-zero this is a new insert
		resp.Message = err.Error()
		return  encode(resp)
	}
	if len(l.Name) == 0 && len(l.Name) > 500 {
		err := errors.New("Name should be between 1 and 500 characters")
		handleError(err)
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
	resp.Message = "Updated successfully"
	return encode(resp)
}

func putLocation(id int, data interface{}) {
	fmt.Println("update locations")
}

func deleteLocation(id int) string {
	var resp Response
	resp.Status = "fail"
	var l Location
	l.Id = id
	err := l.delete()
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return  encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Deleted successfully"
	return encode(resp)
}

// create a location
func(l *Location) create() error {
	result, err := db.Exec("INSERT INTO location(name) VALUES (?)", l.Name)
	handleError(err)
	_ = result
	// update cache
	return err
}

// read a location
func readLocation(id int) (Location, error) {
	prepare, err := db.Prepare("SELECT id, name FROM location WHERE id = ?")
	var result Location
	var Name string
	var Id int
	err = prepare.QueryRow(id).Scan(&Id, &Name)
	handleError(err)
	result.Id = Id
	result.Name = Name
	return result, err
}


func readAllLocations(limit int, offset int) ([]Location, error) {
	prepare, err := db.Query("SELECT id, name FROM location LIMIT ?,?", limit, offset)
	var result []Location
	var l Location
	var Name string
	var Id int
	for prepare.Next() {
		err = prepare.Scan(&Id, &Name)
		handleError(err)
		l.Id = Id
		l.Name = Name
		result = append(result, l)
	}
	return result, err
}

// update a location
func (l *Location) update(){
	fmt.Println("Update");
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