# Golang Microservices in Gin GORM Swagger Docker

## Requirements

* Golang 
* Mysql
* Swag cli
* Mockery

## Steps

### Change the config in .env for database and migrate

### Create a migration file

```
go run cmd/migrate/main.go create product_update sql

```

Modify the migration sql once created in migration folder


```sh
make migrate-up
```

###  Generate API document to the ./doc folder using <strong>swag cli</strong>
```sh
 make docs 
```

###  Run service
```sh
make run 
```

###  Run service
```sh
make test 
```