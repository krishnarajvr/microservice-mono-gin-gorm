# Golang Microservices in Gin GORM Swagger Docker

## Requirements

* Golang - 1.14 recommended
* Mysql - 8.0 
* Swag cli - For generating swagger docs
* Mockery  - For generating mock classes for testing

## Steps

### Change the config in .env for database and other configuration


### Create a migration file - if required

```
make migrate-create NAME=create-user

OR

go run cmd/migrate/main.go create create-user sql

```

Modify the migration sql once created in migration folder


```sh
make migrate-up
```

###  Generate API document to the ./doc folder using <strong>swag cli</strong>
```sh
 make doc
```

###  Run service
```sh
make run 
```

###  Generate mock files
```sh
make mock 
```

###  Test service
```sh
make test 
```

### Folder Structure

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
│       ├── handler.go
│       ├── handler_test.go
│       ├── inject.go
│       ├── mocks
│       │   ├── ITenantRepo.go
│       │   └── ITenantService.go
│       ├── model
│       │   └── tenant.go
│       ├── repo
│       │   └── tenant.go
│       ├── routes.go
│       └── service
│           ├── tenant.go
│           └── tenant_test.go
└── README.md
```