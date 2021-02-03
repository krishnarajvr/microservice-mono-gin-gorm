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
- [x] Krakend Gateway integration
- [x] Common error messages
- [x] Share library across service
- [x] CRUD with pagination support
- [x] module support for service
- [ ] Health endpoint
- [ ] Docker
- [ ] Docker Compose
- [ ] Kubernetes

## Overview

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
 cd account-service
```

###  To setup packages locally
```sh
 go mod vendor
```


### Change the config in .env for database and migrate

```sh
make migrate-up

or

go run cmd/migrate/main.go up
```

###  Generate API document to the ./doc folder using <strong>swag cli</strong>
```sh
 make doc
 
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

### Swagger

- Access the url http://localhost:8082/swagger/index.html

## Folder Structure

```sh
├── app
│   ├── app.go  # App Intitialization
│   ├── database
│   │   ├── db
│   │   │   └── db.go
│   │   └── gorm
│   │       └── gorm.go
│   ├── dbs.go # Database Intitialization
│   └── locale
│       └── locale.go
├── cmd  # Starting point for any application
│   ├── api
│   │   └── main.go # Main application start
│   └── migrate
│       └── main.go # Migration start
├── config
│   └── config.go # App configuration. Any environment specific will get from .env
├── go.mod
├── go.sum
├── locale  # Multi language support for the API
│   ├── el-GR
│   │   └── example.yml
│   ├── en-US
│   │   └── example.yml
│   └── zh-CN
│       └── example.yml
├── log  # Log folder
│   └── micro.log
├── Makefile  # Common command 
├── migration # Migration files
│   ├── 20190805170000_create_user_table.sql
│   └── 20210130204915_create_tenant_table.sql
├── module  # Module contain actual business logic
│   ├── module.go
│   └── tenant  # Tenant Module
│   │   ├── handler.go
│   │   ├── handler_test.go
│   │   ├── inject.go
│   │   ├── mocks
│   │   │   ├── ITenantRepo.go
│   │   │   └── ITenantService.go
│   │   ├── model
│   │   │   └── tenant.go
│   │   ├── repo
│   │   │   └── tenant.go
│   │   ├── routes.go
│   │   └── service
│   │       ├── tenant_service.go
│   │       └── tenant_service_test.go
│   └── user # Tenant Module
│       ├── handler.go
│       ├── handler_test.go
│       ├── inject.go
│       ├── mocks
│       │   ├── IUserRepo.go
│       │   └── IUserService.go
│       ├── model
│       │   └── user.go
│       ├── repo
│       │   └── user.go
│       ├── routes.go
│       └── service
│           ├── user_service.go
│           └── user_service_test.go
└── README.md
```