# zapposapi [![Build Status](https://travis-ci.org/akshaysuresh4793/zapposapi.svg?branch=master)](https://travis-ci.org/akshaysuresh4793/zapposapi)
Zappos.com Restaurant API

# Introduction

This is an API architecture for Restaurants.

The technologies used are as follows:

- **Go** - API development

- **MySQL** - Database

- **Docker and Docker Compose** - Containerization and Container orchestration

- **Redis** - Caching Layer

# Proposed Architecture

![architecture](https://github.com/akshaysuresh4793/zapposapi/blob/master/resources/architecture.jpg "Architecture")

# Code structure

![code structure](https://github.com/akshaysuresh4793/zapposapi/blob/master/resources/code%20structure.jpg "Code structure")

Entities involved are as follows:

- Location

- Restaurant

- Menu

- MenuItem

The relationships are as follows:

- A given location has many restaurants (one-to-many)

- A given restautant has many menus (one-to-many)

- A given menu has many menu-items (one-to-many)

Code structure will come here

# API Documentation

The API documentation URL will come here

Docker image can be downloaded here

# Running the code
- Using Docker
	1. Clone this repository

	2. Run the following command to obtain the Docker images

		`docker-compose build`

	3. Run the following command to start each service

		`docker-compose up`
	
	If you are using docker-machine (Windows and Docker) the base url would be http://192.168.99.100/

	Otherwise the base url would be http://localhost/

# Routes
1. Add new location

`curl -X POST http://{host}/location -H "Content-Type: application/json" -d '{"name": "Boston"}'`

2. Get all locations

`curl http://{host}/locations`

3. Get a particular location

`curl http://{host}/location/{location-id}`

4. Update a location

`curl -X PUT http://{host}/location/{location-id} -H "Content-Type: application/json" -d '{"name": "New Boston"}'`

5. Delete a location

`curl -X DELETE http://{host}/location/{location-id}`

6. Add a new restaurant

`curl -X POST http://{host}/restaurant -H "Content-Type: application/json" -d '{"name": "foo hotel","locationId":"1"}'`

7. Get all restaurants

`curl http://{host}/restaurants`

8. Get a particular restaurant

`curl http://{host}/restaurant/{restaurant-id}`

9. Update a restaurant

`curl -X PUT http://{host}/restaurant/{id} -H "Content-Type: application/json" -d '{"name": "newfoo"}'`

10. Delete a restaurant

`curl -X DELETE http://{host}/restaurant/{restaurant-id}`

11. Add a new menu to a restaurant

`curl -X POST http://{host}/restaurant/{restaurant-id}/menu -H "Content-Type: application/json" -d '{"name": "Breakfast"}'`

12. Get all menus

`curl http://{host}/restaurant/{restaurant-id}/menus`

13. Get a particular menu

`curl http://{host}/restaurant/{restaurant-id}/menu/{menu-id}`

14. Update a menu

`curl http://{host}/restaurant/{restaurant-id}/menu/{menu-id} -H "Content-Type: application/json" -d '{"name": "Lunch"}'`

15. Delete a menu

`curl -X DELETE http://{host}/restaurant/{restaurant-id}/menu/{menu-id}`

16. Add a new menu item to a menu

`curl -X POST http://{host}/restaurant/{restaurant-id}/menu/{menu-id}/menu-item -H "Content-Type: application/json" -d '{"name": "egg", "description":"egg in a basket"}'`

17. Get all menu items

`curl http://{host}/restaurant/{restaurant-id}/menu/{menu-id}/menu-items`

18. Get a particular menu item

`curl http://{host}/restaurant/{restaurant-id}/menu/{menu-id}/menu-item/{menu-item-id}`

19. Update a menu item

`curl http://{host}/restaurant/{restaurant-id}/menu/{menu-id}/menu-item/{menu-item-id} -H "Content-Type: application/json" -d '{"name": "new egg in a basket"}'`

20. Delete a menu item

`curl -X DELETE http://{host}/restaurant/{restaurant-id}/menu/{menu-id}/menu-item/{menu-item-id}`
