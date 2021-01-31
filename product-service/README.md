# Golang Microservices in Gin GORM Swagger Docker

## Requirements

* Golang - 1.14 recommended
* Mysql - 8.0 
* Swag cli - For generating swagger docs
* Mockery  - For generating mock classes for testing

## Steps

### Change the config in .env for database and other configuration



### Create a migration file

```
make migrate-create NAME=user_create

or 

go run cmd/migrate/main.go create user_create sql

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