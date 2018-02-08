package main

import (
	"testing"
)

func TestPostMenu(t *testing.T) {
	connect("zappos_test")
	var menu Menu
	menu.Name = "foobar menu"
	menu.RestaurantId = 2
	output := postMenu(1, menu)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Inserted successfully\",\"Data\":null}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}

func TestGetMenuById(t *testing.T) {
	connect("zappos_test")
	output := getMenuById(1,1)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Data retrieved\",\"Data\":{\"id\":\"1\",\"name\":\"Breakfast\",\"restaurantId\":\"1\"}}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}

func TestDeleteMenu(t *testing.T) {
	connect("zappos_test")
	output := deleteMenu(2, 3)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Deleted successfully\",\"Data\":null}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}
