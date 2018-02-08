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

INSERT INTO location(name) VALUES ('Boston');
INSERT INTO location(name) VALUES ('New York City');

INSERT INTO restaurant(name, location_id) VALUES ('foo', 1);
INSERT INTO restaurant(name, location_id) VALUES ('bar', 2);

INSERT INTO menu(name, restaurant_id) VALUES('Breakfast', 1);
INSERT INTO menu(name, restaurant_id) VALUES('Lunch', 1);

INSERT INTO menu(name, restaurant_id) VALUES('Breakfast', 2);
INSERT INTO menu(name, restaurant_id) VALUES('Lunch', 2);

INSERT INTO menu_item(name, description, restaurant_id, menu_id) VALUES('egg', 'egg in a basket', 1, 1);
INSERT INTO menu_item(name, description, restaurant_id, menu_id) VALUES('salad', 'salad', 1, 2);

INSERT INTO menu_item(name, description, restaurant_id, menu_id) VALUES('egg', 'egg in a basket', 2, 3);
INSERT INTO menu_item(name, description, restaurant_id, menu_id) VALUES('salad', 'salad', 2, 4);