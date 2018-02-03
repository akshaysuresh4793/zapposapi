package main

import(
	"fmt"
)

type Restaurant struct {
	id int
	name string
	location string
}

func getRestaurants() string {
	return "Get Restaurants"
}

func getRestaurantById(id int) string {
	return "Get Restaurant by ID"
}

func postRestaurant(data interface{}) string {
	return "Post Restaurant"
}

func putRestaurant(id int, data interface{}) {
	fmt.Println("update restaurant ", id)
}

func deleteRestaurant(id int) {
	fmt.Println("delete restaurant ", id)
}

// create a restaurant
func(r *Restaurant) create() {
	fmt.Println("Create");
}

// read a restaurant
func (r *Restaurant) read(){
	fmt.Println("Read");
}

// update a restaurant
func (r *Restaurant) update(){
	fmt.Println("Update");
}

// delete a restaurant
func (r *Restaurant) delete() {
	fmt.Println("Delete");
}