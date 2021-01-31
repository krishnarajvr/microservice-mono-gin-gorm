.PHONY: keypair migrate-create migrate-up migrate-down migrate-force

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
	# another - migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)
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
	go test -v ./module/tenant/service
	go test -v ./module/tenant
mock:
	mockery --dir=module/tenant/service --name=ITenantService  --output=module/tenant/mocks
	mockery --dir=module/tenant/repo --name=ITenantRepo  --output=module/tenant/mocks