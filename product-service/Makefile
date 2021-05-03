.PHONY: doc run test mock load-env test-tenant

SERVICE_NAME=Product-Service

export SERVICE_NAME

doc:
	@echo "---Generate doc files---"
	swag init -g module/module.go -o doc 

doc-dep:
	@echo "---Generate doc files---"
	swag init -g cmd/api/main.go -o doc  

migrate-create:
	@echo "---Creating migration files---"
	go run cmd/migrate/main.go create $(NAME) sql

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

migrate-force:
	go run cmd/migrate/main.go force $(VERSION)

run:
	go run cmd/api/main.go 

load-env:
  	export $(cat .env | xargs) && echo $DEBUG
	  
test:
	go test -v ./module/*/*/

test-cover:
	go test -v ./module/product/service -cover
	go test -v ./module/product -cover

test-vendor:	
	go test -v ./module/product/service
	go test -v ./module/product

test-participant:	
	go test -v ./module/participant/service
	go test -v ./module/participant

mock:
	mockery --dir=module/product/service --name=IProductService  --output=module/product/mocks
	mockery --dir=module/product/repo --name=IProductRepo  --output=module/product/mocks
	