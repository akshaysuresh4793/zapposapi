package main

import (
	"fmt"
	"net/http"
	"regexp"
)

// registered methods
var getRestaurantsURL = regexp.MustCompile(`/restaurants`)
var getRestaurantByIdURL = regexp.MustCompile(`/restaurant/\d`)
var getMenusURL = regexp.MustCompile(`/restaurant/\d/menus`)
var getMenuByIdURL = regexp.MustCompile(`/restaurant/\d/menu/\d`)
var getMenuItemsURL = regexp.MustCompile(`/restaurant/\d/menu/\d/menu-items`)
var getMenuItemByIdURL = regexp.MustCompile(`/restaurant/\d/menu/\d/menu-item/\d`)
var getLocationURL = regexp.MustCompile(`/locations`)
var getLocationByIdURL = regexp.MustCompile(`/location/\d`)

var postRestaurantMenuItemURL = regexp.MustCompile(`/restaurant/\d/menu/\d/menu-item`)
var postRestaurantMenuURL = regexp.MustCompile(`/restaurant/\d/menu`)
var postRestaurantURL = regexp.MustCompile(`/restaurant`)
/*
	GET / -> welcome page
	GET /restaurants -> get restaurants list
	GET /restaurant/restaurant-id -> get entire restaurant object matching id
	GET /restaurant/restaurant-id/menus -> get menus matching restaurant id
	GET /restaurant/restaurant-id/menu/menu-id -> get menus matching restaurant id
	GET /restaurant/restaurant-id/menu/menu-id/menu-items -> get menu items matching menus
	GET /restaurant/restaurant-id/menu/menu-id/menu-item/menu-item-id -> get menu item matching matching the menu


	POST /restaurant -> create a new restaurant
	POST /restaurant/restaurant-id/menu -> add a new menu to a restaurant
	POST /restaurant/restaurant-id/menu-id/menu-item -> add a new menu item to a menu

	PUT /restaurant -> update a restaurant attribute
	PUT /restaurant/restaurant-id/menu -> update restaurant menu
	PUT /restaurant/restaurant-id/menu-id/menu-item -> update restaurant menu item

	DELETE /restaurant/restaurant-id -> delete a restaurant
	DELETE /restaurant/restaurant-id/menu-id -> delete a menu item from a restaurant
	DELETE /restaurant/restaurant-id/menu-id/menu-item-id -> delete menu item from a menu


	GET /locations -> get locations
	POST /location -> add a new location
	PUT /location -> update a new location
	DELETE /location -> delete a location

	any other page
	// custom 404 page
*/


func main() {
	// entry point
	fmt.Println("Entry point")
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
	    	case (getRestaurantsURL.MatchString(r.URL.Path) && r.Method == "GET"):
	        	fmt.Fprintf(w, getRestaurants())
	        case (getMenuItemByIdURL.MatchString(r.URL.Path) && r.Method == "GET"):
	        	fmt.Fprintf(w, getMenuItemById(1,1,1))
	        case (getMenuItemsURL.MatchString(r.URL.Path) && r.Method == "GET"):
	        	fmt.Fprintf(w, getMenuItems(1,1))
	         case (getMenuByIdURL.MatchString(r.URL.Path) && r.Method == "GET"):
	        	fmt.Fprintf(w, getMenuById(1,1))
	        case (getMenusURL.MatchString(r.URL.Path) && r.Method == "GET"):
	        	fmt.Fprintf(w, getMenus(1))
	        case (getRestaurantByIdURL.MatchString(r.URL.Path) && r.Method == "GET"):
	        	fmt.Fprintf(w, getRestaurantById(1))
	        case (getLocationURL.MatchString(r.URL.Path) && r.Method == "GET"):
	        	fmt.Fprintf(w, getRestaurantById(1))
	        case (postRestaurantMenuItemURL.MatchString(r.URL.Path) && r.Method == "POST"):
	        	fmt.Fprintf(w, postMenuItem(1,1,1))
	        case (postRestaurantMenuURL.MatchString(r.URL.Path) && r.Method == "POST"):
	        	fmt.Fprintf(w, postMenu(1,1))
	        case (postRestaurantURL.MatchString(r.URL.Path) && r.Method == "POST"):
	        	fmt.Fprintf(w, postRestaurant(1))
	    	default:
	        	fmt.Fprintf(w, "Unknown Pattern")
	    	}
	})
	http.ListenAndServe(":8080", nil)
}

