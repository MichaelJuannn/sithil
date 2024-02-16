# Sithil

This is a Go Fiber project that includes various services such as cart management, order processing, product management, and user management.

## Prequisites

- Go 1.21 or later
- MySQL 8 or later
- docker (optional)

## Database Implementation

The project uses GORM, a Go ORM, for database operations. The database schema includes tables for categories, products, users, carts, and orders. 

The `database/connect.go` file sets up the database connection, and the `database/seed.go` file seeds the database with initial data.

The `model` package defines the structs for the different tables, and GORM uses these structs to map between the Go code and the database.

## RESTful API Design

The project provides a RESTful API for managing users, products, and carts. The API follows standard REST conventions:

### User API Routes

- `GET /api/users/` - Protected route, tests JWT authentication.
- `POST /api/users/register` - Register a new user.
- `POST /api/users/login` - Login a user.

### Product API Routes

- `GET /api/products/` - Get all products.

### Cart API Routes

- `POST /api/carts/` - Protected route, add a product to the cart.
- `GET /api/carts/` - Protected route, get all products in the cart.
- `DELETE /api/carts/` - Protected route, delete a product from the cart.
- `GET /api/carts/checkout` - Protected route, checkout the cart.

The `middleware.Protected()` function is used to protect routes that require the user to be authenticated. The user must send a valid JWT in the `Authorization` header of their request to access these routes.

## Getting Started

### Using Docker

If you have Docker installed, you can run the project in a Docker container:

1. Switch to the Docker branch:

```sh
git checkout docker_compose
```
2. Run Docker Compose :
    
```sh
docker-compose up
```


### Without Docker
1. Clone the repository:

```sh
git clone https://github.com/MichaelJuannn/sithil.git
```

2. Navigate to the project directory:

```sh
cd sithil
```

3. Install the project dependencies:

```sh
go mod download
```

4. Run the project:

```sh
go run main.go
```
dont forget to set the environment variable in .env file
here is the example 
```sh
DB_HOST=db_host
DB_NAME=sithil
DB_USER=user
DB_PASSWORD=password
DB_PASSWORD_ROOT=example
DB_PORT=3306
JWT_SECRET=xdd
```


The project will be available at `http://localhost:8000`.