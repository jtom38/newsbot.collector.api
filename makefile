.PHONY: help
help: ## Shows this help command
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## builds the application with the current go runtime
	sqlc generate
	~/go/bin/swag f
	~/go/bin/swag i
	go build .
	
docker-build: ## Generates the docker image
	docker build -t "newsbot.collector.api" .
	docker image ls | grep newsbot.collector.api

migrate-dev: ## Apply sql migrations to dev db
	goose -dir "./database/migrations" postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up 

migrate-dev-down: ## revert sql migrations to dev db
	goose -dir "./database/migrations" postgres "user=postgres password=postgres dbname=postgres sslmode=disable" down 

swag: ## Generates the swagger documentation with the swag tool
	~/go/bin/swag f
	~/go/bin/swag i

gensql: ## Generates SQL code with sqlc
	sqlc generate