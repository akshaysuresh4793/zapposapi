package main

import (
	"testing"
)

func TestPostRestaurant(t *testing.T) {
	connect("zappos_test")
	var restaurant Restaurant
	restaurant.Name = "foobar hotel"
	restaurant.LocationId = 1
	output := postRestaurant(restaurant)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Inserted successfully\",\"Data\":null}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}

func TestGetRestaurantById(t *testing.T) {
	connect("zappos_test")
	output := getRestaurantById(1)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Data retrieved\",\"Data\":{\"id\":\"1\",\"name\":\"foo\",\"locationId\":\"1\"}}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}

func TestDeleteRestaurant(t *testing.T) {
	connect("zappos_test")
	output := deleteRestaurant(1)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Deleted successfully\",\"Data\":null}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}
