# Account Service

## Requirements

* Golang - 1.14 recommended
* Mysql - 8.0 
* [Swag cli](https://github.com/swaggo/swag)  - For generating swagger docs
* [Mockery](https://github.com/vektra/mockery) - For generating mock classes for testing


- Note: Once Swag and Mockery installed ```swag``` and ```mockery``` command should work in terminal. 
- Tip: can copy the build binary of the package to {home}/go/bin path also works. Also can change the Makefile command path as well.

## Steps

### Coniguration
Change the config in .env for database and other configuration


### Database Migration

Existing migrations

```sh
make migrate-up
```
Create a migration file - if required

```
make migrate-create NAME=create-product
```
Modify the migration sql(create-product-1123213.sql) once created in migration folder and then run again ```make migrate-up```.


### Swagger Document  

Generate API document to the ./doc folder using <strong>swag cli</strong>
```sh
 make doc
```

###  Run service

```sh
make run 
```

Make sure you run ```make migrate-up``` before the running

### Testing

Generate mock files
```sh
make mock 
```

Test service
```sh
make test 
```

