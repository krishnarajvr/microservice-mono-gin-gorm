.PHONY: doc run test mock load-env test-tenant

GATEWAY_URL 	=http://localhost:8080
ACCOUNT_PREFIX	=/account/v1

doc:
	@echo "---Generate doc files---"
	swag init -g module/module.go -o doc

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

load-env:
  	export $(cat .env | xargs) && echo $(DEBUG)

test:
	go test -v ./module/*/*/

test-tenant:	
	go test -v ./module/tenant/service
	go test -v ./module/tenant	
	
test-user:	
	go test -v ./module/user/service
	go test -v ./module/user

test-client:	
	go test -v ./module/client/service -cover
	go test -v ./module/client -cover

mock:
	mockery --dir=module/tenant/service --name=ITenantService  --output=module/tenant/mocks
	mockery --dir=module/tenant/repo --name=ITenantRepo  --output=module/tenant/mocks
	mockery --dir=module/user/service --name=IUserService  --output=module/user/mocks
	mockery --dir=module/user/repo --name=IUserRepo  --output=module/user/mocks	
	mockery --dir=module/client/service --name=IClientService  --output=module/client/mocks
	mockery --dir=module/client/repo --name=IClientRepo  --output=module/client/mocks	