CREATE DATABASE IF NOT EXISTS zappos;
USE zappos;

-- location
CREATE TABLE location(
	id int PRIMARY KEY AUTOINCREMENT,
	name varchar2(500)
)

-- restaurant
CREATE TABLE restaurant(
	id int PRIMARY KEY AUTOINCREMENT,
	name varchar2(500),
	location_id REFERENCES location(id)
)

-- menu
CREATE TABLE menu(
	id int PRIMARY KEY AUTOINCREMENT,
	name varchar2(500),
	restaurant_id REFERENCES restaurant(id)
)

-- menu_item
CREATE TABLE menu_item(
	id int PRIMARY KEY AUTOINCREMENT,
	name varchar2(500),
	menu_id REFERENCES menu(id)
)