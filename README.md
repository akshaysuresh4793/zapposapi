# zapposapi [![Build Status](https://travis-ci.org/akshaysuresh4793/zapposapi.svg?branch=master)](https://travis-ci.org/akshaysuresh4793/zapposapi)
Zappos.com Restaurant API

# Introduction

This is an API architecture for Restaurants.

The technologies used are as follows:

- **Go** - API development

- **MySQL** - Database

- **Docker** - Containerization

- **Redis** - Caching Layer

# Proposed Architecture

![architecture](https://github.com/akshaysuresh4793/zapposapi/blob/master/resources/code%20structure.jpg "Architecture")

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
