package main

import (
	"testing"
)

func TestPostLocation(t *testing.T) {
	connect("zappos_test")
	var location Location
	location.Name = "foo"
	output := postLocation(location)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Inserted successfully\",\"Data\":null}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}

func TestGetLocationById(t *testing.T) {
	connect("zappos_test")
	output := getLocationById(1)
	expectedOutput := "{\"Status\":\"success\",\"Message\":\"Data retrieved\",\"Data\":{\"Id\":1,\"name\":\"foo\"}}"
	if output != expectedOutput {
		t.Errorf("Mismatch - Expected : " + expectedOutput + ", output : " + output)
	}
}
