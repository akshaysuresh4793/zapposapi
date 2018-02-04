CREATE DATABASE IF NOT EXISTS zappos;
USE zappos;

-- location
CREATE TABLE IF NOT EXISTS location(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500));

-- restaurant
CREATE TABLE IF NOT EXISTS restaurant(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	location_id int REFERENCES location(id));

-- menu
CREATE TABLE IF NOT EXISTS menu(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	restaurant_id int REFERENCES restaurant(id));

-- menu_item
CREATE TABLE IF NOT EXISTS menu_item(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	menu_id int REFERENCES menu(id)
);