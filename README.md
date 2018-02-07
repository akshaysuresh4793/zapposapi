# zapposapi
Zappos.com Restaurant API

# Introduction

This is an API architecture for Restaurants.

The technologies used are as follows:

- **Go** - API development

- **MySQL** - Database

- **Docker** - Containerization

- **Redis** - Caching Layer

- **Supervisor** - Run the API as child processes so that it starts again if it terminates with non-zero exit code

# Proposed Architecture

Architecture diagram will come here

# Code structure

The code structure can be illustrated by the following diagram:


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
