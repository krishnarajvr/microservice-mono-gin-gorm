# Golang Microservices in Gin GORM Swagger Docker

## Requirements

* Golang - 1.14 recommended
* Mysql - 8.0  
* [Swag cli](https://github.com/swaggo/swag)  - For generating swagger docs
* [Mockery](https://github.com/vektra/mockery) - For generating mock classes for testing

## Featues

- [x] Config using .env
- [x] Database migration 
- [x] GORM Integration
- [x] Dependency Injection
- [x] Swagger docs
- [x] Repository
- [x] Seperate service
- [x] Seperate handler
- [x] Mono Repo
- [x] Multi language
- [x] Json logger
- [x] Makefile for commands
- [x] Mock object integration 
- [x] Unit Test
- [x] Integration Test
- [x] Standard request and response and errors
- [x] Common form validation 
- [x] Health endpoint
- [x] Krakend Gateway integration
- [ ] Common error messages
- [ ] Docker
- [ ] Docker Compose
- [ ] Kubernetes
- [ ] Share library across service
- [x] CRUD with pagination support

Overview

![image Architecture](https://raw.githubusercontent.com/krishnarajvr/microservice-mono-gin-gorm/master/assets/golang-monorepo.png) 

#### Gateway
Which act as a singe entrypoint for all the services
- Handle logs,trace,metrics collection
- Handle Authentication
- Handle ratelimit, Circuit breaker and many more

#### Common Library
- Common functionality shared accross the microserices

###  Go to specific microservice
```sh
 cd product-service
```
### Change the config in .env for database and migrate

```sh
make migrate-up

or

go run cmd/migrate/main.go up
```

###  Generate API document to the ./doc folder using <strong>swag cli</strong>
```sh
 make docs
 
 or
 
 swag init -g cmd/api/main.go -o doc
```

Swagger docs will be available at /swagger/index.html of the api path

###  Run service
```sh
make run 

or 

go run cmd/api/main.go
```

###  Generate mock files
```sh
make mock 
```

###  Test service
```sh
make test 
```

## Folder Structure

```sh
├── product-service
├── README.md
├── .env  # Environment configuration
└── user-service
    ├── app  # App specific folders
    │   ├── database
    │   │   ├── adapter
    │   │   │   ├── db
    │   │   │   │   └── db.go
    │   │   │   └── gorm
    │   │   │       └── gorm.go
    │   │   ├── database.go
    │   │   └── inject.go
    │   ├── locale
    │   │   ├── inject.go
    │   │   └── locale.go
    │   └── middleware
    │       └── logger.go
    ├── cmd  # Any Command line running functions
    │   ├── api
    │   │   └── main.go  # Starting point for the applicatioin
    │   └── migrate
    │       └── main.go  # Migrate util for db
    ├── doc  # Swagger docs auto generate available at /swagger/index.html
    │   ├── docs.go
    │   ├── swagdto
    │   │   └── user.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── common # Common modules
    │   ├── common.go 
    │   ├── error.go # Error responses
    │   ├── middleware 
    │   ├── response.go # Response structure
    │   ├── util 
    │   └── validation  # Validation utils
    ├── config
    │   └── config.go
    ├── locale  # Locale language support
    │   ├── el-GR
    │   │   └── example.yml
    │   ├── en-US
    │   │   └── example.yml
    │   └── zh-CN
    │       └── example.yml
    ├── log
    │   ├── micro.log -> micro.log.20201208.log
    │   └── micro.log.20201208.log
    ├── go.mod
    ├── go.sum
    ├── migration # Migration scripts
    │   └── 20190805170000_create_user_table.sql
    ├── model   # Contain model and dtos
    │   └── user.go
    ├── repo
    │   ├── mocks # Mock objects
    │   │   └── IUserRepo.go
    │   └── user.go
    ├── handler # define the rest endpoints
    │   ├── user.go
    │   └── handler.go
    └── service # Services do the business logic
        └── user.go

```