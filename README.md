# Golang Microservices using Gin GORM Swagger Docker

## Requirements

* Golang - 1.5 or higher recommended
* Mysql - 8.0  or higher
* [Swag cli](https://github.com/swaggo/swag)  - For generating swagger docs
* [Mockery](https://github.com/vektra/mockery) - For generating mock classes for testing
* [Docker](https://docs.docker.com/engine/install/) and [Docker-compose](https://docs.docker.com/compose/install/) for the containerized setup - Not mandatory

Install Swag and Mockery

```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/vektra/mockery/cmd/mockery

```
## Featues

- [x] Multi tenancy support
- [x] Scalable folder structure
- [x] Config using .env
- [x] Database migration 
- [x] GORM Integration
- [x] Dependency Injection
- [x] Swagger docs 
- [x] Separate handler, service, repository(repo), model
- [x] Multi language support
- [x] Json logger 
- [x] Makefile for commands
- [x] Mock object integration 
- [x] Unit Test
- [x] Integration Test
- [x] Standard request and response and errors
- [x] Common form validation 
- [ ] Health endpoint
- [x] Krakend Gateway integration
- [x] Common error messages
- [x] Docker
- [x] Docker Compose
- [x] Share library across service
- [x] CRUD with pagination support
- [ ] Kubernetes

## Overview

![image Architecture](https://raw.githubusercontent.com/krishnarajvr/microservice-mono-gin-gorm/master/assets/golang-monorepo.png) 

#### Gateway
Which act as a singe entrypoint for all the services
- Handle logs,trace,metrics collection
- Handle Authentication
- Handle ratelimit, Circuit breaker and many more

#### Common Library
- Common functionality shared accross the microserices , Refer [micro-common](https://github.com/krishnarajvr/micro-common)

#### Account Service
- Manage tenants, users, authentication and authrorization

#### Product Service
- Manage product catalog. Demo purpose


#### Running Application

Go to the folders ```./gateway``` ```./account-service``` ```./product-service``` and follow the readme.md files

###  Eg: Build Account microservice
```sh
 cd account-service
```

###  Setup packages locally
```sh
 go mod vendor
```

### Change the config in .env for database and migrate the tables
```sh
make migrate-up
```

###  Generate API document to the ./doc folder using <strong>swag cli</strong>
```sh
 make doc
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
- or through gateway once krakend is started http://localhost:8080/account/swagger/index.html 

Same steps can be followed for product-service.  For gateway steps are different. Please refere readme.md for gateway

### Useful go commands
```sh
go clean --modcache   # Clean the mod cache
go mod vendor   # Initilize vendor folder locally
go mod download  #Download missing packages
```

###  Docker Compose setup
If Docker and Docker compose installed. Run below command from root directory

```sh
sudo docker-compose up
```

### Create an account that can be used for /adminLogin api for token generation. Refer account service swagger.
```
curl -X POST "https://localhost:8080/account/v1/tenantRegister" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"domain\": \"eBook\",  \"email\": \"tenant1@mail.com\",  \"firstName\": \"John\",  \"lastName\": \"Doe\",  \"name\": \"Tenant1\",  \"password\": \"Pass@1\"}"
```

Can use the same user email and password for /adminLogin api

## Folder Structure
```sh
├── app  # App Intitialization
│   ├── app.go
│   └── database
│       ├── db
│       │   └── db.go
│       └── gorm
│           └── gorm.go
├── cmd # Starting point for any application
│   ├── api
│   │   └── main.go # Main application start
│   └── migrate
│       └── main.go # Migration start
├── config
│   └── config.go # App configurations
├── deploy
│   ├── bin
│   │   └── init.sh
│   ├── Dockerfile
│   └── prod.Dockerfile  
├── doc  # Swagger doc - Autogenerated
│   ├── docs.go
│   ├── swagdto
│   │   └── error.go  # Common errors used for swagger - Custom
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── locale  # Language files
│   ├── el-GR
│   │   └── language.yml
│   ├── en-US
│   │   └── language.yml
│   └── zh-CN
│       └── language.yml
├── log
│   ├── micro.log -> micro.log.20210216.log
│   └── micro.log.20210216.log
├── Makefile
├── migration  # Migration files
│   └── 20210316142300_create_product.sql
├── module  # Application module - Main buisiness logic
│   ├── module.go
│   └── product
│       ├── handler.go
│       ├── handler_test.go
│       ├── inject.go
│       ├── mocks
│       │   ├── IProductRepo.go
│       │   └── IProductService.go
│       ├── model
│       │   └── product.go
│       ├── repo
│       │   └── product.go
│       ├── routes.go
│       ├── service
│       │   ├── product_service.go
│       │   └── product_service_test.go
│       └── swagger
│           └── product.go
├── README.md
└── util
```
