package main
import (
	"fmt"
	"net/http"
	mu "github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

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
	// connect to db
	_ = connect()


	// entry point
	mux := mu.NewRouter()



	/* ------------ Restaurant ------------- */

	mux.HandleFunc("/restaurant",  func(w http.ResponseWriter, r *http.Request) {
			data := handleData(r,"restaurant")
			result := postRestaurant(data.(Restaurant))
			fmt.Fprintf(w, result)
	}).Methods("POST")

	mux.HandleFunc("/restaurants",  func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			start := 0
			limit := 10
			var err error
			if q["start"] != nil {
				start,err = strconv.Atoi(q["start"][0])
				handleError(err)
			}
			if q["limit"] != nil {
				limit,err = strconv.Atoi(q["limit"][0])
				handleError(err)
			}
			result := getRestaurants(start, limit)
			fmt.Fprintf(w, result)}).Methods("GET")

	mux.HandleFunc("/restaurant/{restaurant-id}",  func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			restaurantId, err := strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			result := getRestaurantById(restaurantId)
			fmt.Fprintf(w, result)
	}).Methods("GET")

	mux.HandleFunc("/restaurant/{restaurant-id}",  func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			data := handleData(r,"raw")
			_ = data
				restaurantId, err := strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			result := putRestaurant(restaurantId, data.(map[string]interface{}))
			fmt.Fprintf(w, result)
	}).Methods("PUT")

	mux.HandleFunc("/restaurant/{restaurant-id}", func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			restaurantId, err := strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			result := deleteRestaurant(restaurantId)
			fmt.Fprintf(w, result)
	}).Methods("DELETE")

	/* --------------- Menu --------------- */

	mux.HandleFunc("/restaurant/{restaurant-id:[0-9]+}/menu",  func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			data := handleData(r,"menu")
			restaurantId, err := strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			result := postMenu(restaurantId, data.(Menu))
			fmt.Fprintf(w, result)
	}).Methods("POST")

	mux.HandleFunc("/restaurant/{restaurant-id}/menus",  func(w http.ResponseWriter, r *http.Request) {
			var err error
			var restaurantId int
			vars := mu.Vars(r)
			restaurantId, err = strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			q := r.URL.Query()
			start := 0
			limit := 10
			var result string
			if q["start"] != nil {
				start,err = strconv.Atoi(q["start"][0])
				handleError(err)
			}
			if q["limit"] != nil {
				limit,err = strconv.Atoi(q["limit"][0])
				handleError(err)
			}
			result = getMenus(restaurantId, start, limit)
			// handleError(err)
			fmt.Fprintf(w, result)
	}).Methods("GET")

	mux.HandleFunc("/restaurant/{restaurant-id}/menu/{menu-id}",  func(w http.ResponseWriter, r *http.Request) {
			var err error
			var restaurantId int
			var menuId int
			var result string
			vars := mu.Vars(r)
			restaurantId, err = strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err = strconv.Atoi(vars["menu-id"])
			handleError(err)
			result = getMenuById(restaurantId, menuId)
			// handleError(err)
			fmt.Fprintf(w, result)
	}).Methods("GET")

	mux.HandleFunc("/restaurant/{restaurant-id}/menu/{menu-id}",  func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			data := handleData(r,"raw")
			_ = data
			restaurantId, err := strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err := strconv.Atoi(vars["menu-id"])
			handleError(err)
			result := putMenu(restaurantId, menuId, data.(map[string]interface{}))
			fmt.Fprintf(w, result)
	}).Methods("PUT")

	mux.HandleFunc("/restaurant/{restaurant-id}/menu/{menu-id}", func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			var err error
			var restaurantId int
			var menuId int
			restaurantId, err = strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err = strconv.Atoi(vars["menu-id"])
			handleError(err)
			result := deleteMenu(restaurantId, menuId)
			fmt.Fprintf(w, result)
	}).Methods("DELETE")

	/* -------------------- Menu Items ---------------------- */

	mux.HandleFunc("/restaurant/{restaurant-id:[0-9]+}/menu/{menu-id:[0-9]+}/menu-item",  func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			data := handleData(r,"menu-item")
			restaurantId, err := strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err := strconv.Atoi(vars["menu-id"])
			handleError(err)
			result := postMenuItem(restaurantId, menuId, data.(MenuItem))
			fmt.Fprintf(w, result)
	}).Methods("POST")

	mux.HandleFunc("/restaurant/{restaurant-id}/menu/{menu-id}/menu-items",  func(w http.ResponseWriter, r *http.Request) {
			var err error
			var restaurantId int
			var menuId int
			var result string
			q := r.URL.Query()
			start := 0
			limit := 10
			vars := mu.Vars(r)
			restaurantId, err = strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err = strconv.Atoi(vars["menu-id"])
			handleError(err)
			if q["start"] != nil {
				start,err = strconv.Atoi(q["start"][0])
				handleError(err)
			}
			if q["limit"] != nil {
				limit,err = strconv.Atoi(q["limit"][0])
				handleError(err)
			}
			result = getMenuItems(restaurantId, menuId, start, limit)
			fmt.Fprintf(w, result)
	}).Methods("GET")

	mux.HandleFunc("/restaurant/{restaurant-id}/menu/{menu-id}/menu-item/{menu-item-id}",  func(w http.ResponseWriter, r *http.Request) {
			var err error
			var restaurantId int
			var menuId int
			var menuItemId int
			var result string
			vars := mu.Vars(r)
			restaurantId, err = strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err = strconv.Atoi(vars["menu-id"])
			handleError(err)
			menuItemId, err = strconv.Atoi(vars["menu-item-id"])
			handleError(err)
			result = getMenuItemById(restaurantId, menuId, menuItemId)
			// handleError(err)
			fmt.Fprintf(w, result)
	}).Methods("GET")

	mux.HandleFunc("/restaurant/{restaurant-id}/menu/{menu-id}/menu-item/{menu-item-id}",  func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			data := handleData(r,"raw")
			_ = data
			restaurantId, err := strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err := strconv.Atoi(vars["menu-id"])
			handleError(err)
			menuItemId, err := strconv.Atoi(vars["menu--item-id"])
			handleError(err)
			result := putMenuItem(restaurantId, menuId, menuItemId, data.(map[string]interface{}))
			fmt.Fprintf(w, result)
	}).Methods("PUT")

	mux.HandleFunc("/restaurant/{restaurant-id}/menu/{menu-id}/menu-item/{menu-item-id}", func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			var err error
			var restaurantId int
			var menuId int
			var menuItemId int
			restaurantId, err = strconv.Atoi(vars["restaurant-id"])
			handleError(err)
			menuId, err = strconv.Atoi(vars["menu-id"])
			handleError(err)
			menuItemId, err = strconv.Atoi(vars["menu-item-id"])
			handleError(err)
			result := deleteMenuItem(restaurantId, menuId, menuItemId)
			fmt.Fprintf(w, result)
	}).Methods("DELETE")

	/* -------------------- Location -------------------- */

	mux.HandleFunc("/location",  func(w http.ResponseWriter, r *http.Request) {
			data := handleData(r, "location")
			result := postLocation(data.(Location))
			fmt.Fprintf(w, result)
	}).Methods("POST")

	mux.HandleFunc("/locations",  func(w http.ResponseWriter, r *http.Request) {
			q := r.URL.Query()
			start := 0
			limit := 10
			var err error
			if q["start"] != nil {
				start,err = strconv.Atoi(q["start"][0])
				handleError(err)
			}
			if q["limit"] != nil {
				limit,err = strconv.Atoi(q["limit"][0])
				handleError(err)
			}
			result := getLocations(start, limit)
			fmt.Fprintf(w, result)
	}).Methods("GET")

	mux.HandleFunc("/location/{location-id}",  func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			locationId, err := strconv.Atoi(vars["location-id"])
			handleError(err)
			result :=  getLocationById(locationId)
			fmt.Fprintf(w, result)
	}).Methods("GET")

	mux.HandleFunc("/location/{location-id}", func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			data := handleData(r,"raw")
			_ = data
			locationId, err := strconv.Atoi(vars["location-id"])
			handleError(err)
			result := putLocation(locationId, data.(map[string]interface{}))
			fmt.Fprintf(w, result)
	}).Methods("PUT")

	mux.HandleFunc("/location/{location-id}", func(w http.ResponseWriter, r *http.Request) {
			vars := mu.Vars(r)
			locationId, err := strconv.Atoi(vars["location-id"])
			handleError(err)
			result := deleteLocation(locationId)
			fmt.Fprintf(w, result)
	}).Methods("DELETE")

	/* ------------------ Not Found ------------ */

	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			var resp Response
			resp.Status = "fail"
			resp.Message = "404 Page not found"
			fmt.Fprintf(w, encode(resp))
	})

	mux.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			var resp Response
			resp.Status = "fail"
			resp.Message = "405 Method not allowed"
			fmt.Fprintf(w, encode(resp))
	})

	fmt.Println("Entry point")
	fmt.Println("Listening on port 8080")
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}


func handleData(r *http.Request, entity string) (interface{}) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	handleError(err)
	if err != nil {
		return nil
	}
	switch entity {
	case "restaurant":
		var r Restaurant
		err = json.Unmarshal(b, &r)
		handleError(err)
		return r
	case "menu":
		var m Menu
		err = json.Unmarshal(b, &m)
		handleError(err)
		return m
	case "menu-item":
		var m MenuItem
		err = json.Unmarshal(b, &m)
		handleError(err)
		return m
	case "location":
		var l Location
		err = json.Unmarshal(b, &l)
		handleError(err)
		return l
	default:
		var def map[string]interface{}
		err = json.Unmarshal(b, &def)
		handleError(err)
		return def
	}	
}