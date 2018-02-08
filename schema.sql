CREATE DATABASE IF NOT EXISTS zappos;
USE zappos;

-- location
CREATE TABLE IF NOT EXISTS location(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	INDEX(id));

-- restaurant
CREATE TABLE IF NOT EXISTS restaurant(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	location_id int,
	FOREIGN KEY(location_id) REFERENCES location(id),
	INDEX(id));

-- menu
CREATE TABLE IF NOT EXISTS menu(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	restaurant_id int,
	FOREIGN KEY(restaurant_id) REFERENCES restaurant(id),
	INDEX(id));

-- menu_item
CREATE TABLE IF NOT EXISTS menu_item(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	description varchar(1000),
	restaurant_id int,
	FOREIGN KEY(restaurant_id) REFERENCES restaurant(id),
	menu_id int,
	FOREIGN KEY(menu_id) REFERENCES menu(id),
	INDEX(id));


CREATE DATABASE IF NOT EXISTS zappos_test;
USE zappos_test;

-- location
CREATE TABLE IF NOT EXISTS location(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	INDEX(id));

-- restaurant
CREATE TABLE IF NOT EXISTS restaurant(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	location_id int,
	FOREIGN KEY(location_id) REFERENCES location(id),
	INDEX(id));

-- menu
CREATE TABLE IF NOT EXISTS menu(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	restaurant_id int,
	FOREIGN KEY(restaurant_id) REFERENCES restaurant(id),
	INDEX(id));

-- menu_item
CREATE TABLE IF NOT EXISTS menu_item(
	id int PRIMARY KEY AUTO_INCREMENT,
	name varchar(500),
	description varchar(1000),
	restaurant_id int,
	FOREIGN KEY(restaurant_id) REFERENCES restaurant(id),
	menu_id int,
	FOREIGN KEY(menu_id) REFERENCES menu(id),
	INDEX(id));