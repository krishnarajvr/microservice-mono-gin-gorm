# Golang Microservices in Gin GORM Swagger Docker

## Requirements

* Golang
* Mysql
* Swag cli

## Featues

- [x] Config using .env
- [x] Database migration 
- [x] Swagger docs
- [x] Repository
- [x] Seperate route
- [x] Seperate controller
- [x] Mono Repo
- [x] Multi language
- [x] Docker
- [] Docker Compose
- [] Kubernetes
- [] Unit Test
- [] Integration Test
- [] Mock server integration 
- [] Common validation and errors
- [] Share library across service
- [] CRUD with pagination support
- [] Mongo db integration
- [] Elastic Search integration

###  Go to specific microservice
```sh
 cd user-service
```
### Change the config in .env for database and migrate

```sh
 go run cmd/migrate/main.go up
```

###  Generate API document to the ./doc folder using <strong>swag cli</strong>
```sh
 swag init -g cmd/api/main.go -o doc
```

###  Run service
```sh
 go run cmd/api/main.go
```

## Folder Structure

```sh

├── product-service
├── README.md
├── .env  # Environment configuration
└── user-service
    ├── app  # App specific folders
    │   └── database
    │       ├── adapter
    │       │   ├── db
    │       │   │   └── db.go
    │       │   └── gorm
    │       │       └── gorm.go
    │       ├── database.go
    │       └── inject.go
    ├── cmd  # Any Command line running functions
    │   ├── api
    │   │   └── main.go  # Starting point for the applicatioin
    │   └── migrate
    │       └── main.go  # Migrate util for db
    ├── common # Common modules
    │   ├── common.go 
    │   ├── error.go # Error responses
    │   ├── middleware 
    │   ├── response.go # Response structure
    │   ├── util 
    │   └── validation  # Validation utils
    ├── config
    │   └── config.go
    ├── controller # Handle request, response, input validation and errors
    │   └── user.go
    ├── doc  # Autogenerated Swagger doc using swag
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── dto
    │   └── user.go
    ├── go.mod
    ├── go.sum
    ├── migration # Migration scripts
    │   └── 20190805170000_create_user_table.sql
    ├── model
    │   └── user.go
    ├── repo
    │   └── user.go
    ├── route # define the rest endpoints
    │   ├── route.go
    │   └── v1
    │       └── user.go
    └── service # Services do the business logic
        └── user.go

```