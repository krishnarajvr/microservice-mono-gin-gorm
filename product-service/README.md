# Account Service

## Requirements

* Golang - 1.14 recommended
* Mysql - 8.0 
* [Swag cli](https://github.com/swaggo/swag)  - For generating swagger docs
* [Mockery](https://github.com/vektra/mockery) - For generating mock classes for testing


- Note: Once Swag and Mockery installed ```swag``` and ```mockery``` command should work in terminal. 
- Tip: can copy the build binary of the package to {home}/go/bin path also works. Also can change the Makefile command path as well.

## Steps

### Change the config in .env for database and other configuration


### Create a migration file - if required

```
make migrate-create NAME=create-product

OR

go run cmd/migrate/main.go create create-product sql

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

Make sure you run ```make migrate-up``` before the running

###  Generate mock files
```sh
make mock 
```

###  Test service
```sh
make test 
```

