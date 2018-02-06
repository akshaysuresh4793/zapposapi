package main

import(
	"fmt"
	"errors"
)

type Restaurant struct {
	Id int `json: "id"`
	Name string `json:"name"`
	LocationId int `json:"locationId,string"`
}

func getRestaurants(limit int, offset int) string {
	var resp Response
	resp.Status = "fail"
	result, err := readAllRestaurants(limit, offset)
	handleError(err)
	if(err != nil) {
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
	result, err := readRestaurant(id)
	handleError(err)
	if(err != nil) {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Data retrieved"
	resp.Data = result
	return encode(resp)
}

func postRestaurant(r Restaurant) string {
	var resp Response
	resp.Status = "fail"
	var err error
	if(r.Id != 0) {
		// id cannot be non-zero this is a new insert
		err = errors.New("Id is nonzero")
		handleError(err)
		resp.Message = err.Error()
		return encode(resp)
	}
	if len(r.Name) == 0 && len(r.Name) > 500 {
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
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Inserted successfully"
	return encode(resp)
}

func putRestaurant(id int, data interface{}) {
	fmt.Println("update restaurant ", id)
}

func deleteRestaurant(id int) string {
	var resp Response
	resp.Status = "fail"
	var r Restaurant
	r.Id = id
	err := r.delete()
	handleError(err)
	if err != nil {
		resp.Message = err.Error()
		return encode(resp)
	}
	resp.Status = "success"
	resp.Message = "Deleted successfully"
	return encode(resp)
}

// create a restaurant
func(r *Restaurant) create() (error) {
	// prepare MySQL query
	// insert
	result, err := db.Exec("INSERT INTO restaurant(name, location_id) VALUES (?, ?)", r.Name, r.LocationId)
	handleError(err)
	_ = result
	// update cache
	return err
}

// read a restaurant
func readRestaurant(id int) (Restaurant, error) {
	prepare, err := db.Prepare("SELECT id, name, location_id FROM restaurant WHERE id = ?")
	var result Restaurant
	var Name string
	var Id int
	var Location int
	err = prepare.QueryRow(id).Scan(&Id, &Name, &Location)
	handleError(err)
	result.Id = Id
	result.Name = Name
	result.LocationId = Location
	return result, err
}


func readAllRestaurants(limit int, offset int) ([]Restaurant, error) {
	prepare, err := db.Query("SELECT id, name, location_id FROM restaurant LIMIT ?,?", limit, offset)
	var result []Restaurant
	var r Restaurant
	var Name string
	var Id int
	var Location int
	for prepare.Next() {
		err = prepare.Scan(&Id, &Name, &Location)
		handleError(err)
		r.Id = Id
		r.Name = Name
		r.LocationId = Location
		result = append(result, r)
	}
	return result, err
}

// update a restaurant
func (r *Restaurant) update(){
	fmt.Println("Update");
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