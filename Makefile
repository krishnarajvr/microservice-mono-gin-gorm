.PHONY: compose gateway

compose:
	sudo docker-compose up --build --remove-orphans

gateway:
	sudo docker-compose up --build gateway

account:
	sudo docker-compose up --build account-service