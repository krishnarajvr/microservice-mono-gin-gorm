.PHONY: docs run

# Sample configs
PWD = $(shell pwd)
MPATH = $(PWD)/migrations
PORT = 5432

# Default number of migrations to execute up or down
N = 1

docs:
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

test:
	go test -v ./handler
	go test -v ./service
mock:
	mockery --dir=service --name=IProductService  --output=service/mocks
	mockery --dir=repo --name=IProductRepo  --output=repo/mocks
docker-build:
	sudo docker build -t microapp:v1 -f docker/app/Dockerfile .
docker-run:
	sudo docker start --publish 8092:8080 --name micro-app-v1 microapp:v1
