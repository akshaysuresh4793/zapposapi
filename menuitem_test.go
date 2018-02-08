package main

import (
	"testing"
)

func TestPostMenuItem(t *testing.T) {
	connect("zappos_test")
	var menuItem MenuItem
	menuItem.Name = "foobar menu item"
	menuItem.Description = "This is a foo bar menu item"
	menuItem.RestaurantId = 1
	menuItem.MenuId = 1
	output := postMenuItem(1, 1, menuItem)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Inserted successfully\",\"Data\":null}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}

func TestGetMenuItemById(t *testing.T) {
	connect("zappos_test")
	output := getMenuItemById(1, 1, 1)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Data retrieved\",\"Data\":{\"id\":\"1\",\"name\":\"egg\",\"menuId\":\"1\",\"restaurantId\":\"1\",\"description\":\"egg in a basket\"}}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}

func TestDeleteMenuItem(t *testing.T) {
	connect("zappos_test")
	output := deleteMenuItem(2, 3, 3)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Deleted successfully\",\"Data\":null}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}
