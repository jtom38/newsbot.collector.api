.PHONY: help
help: ## Shows this help command
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## builds the application with the current go runtime
	go build .
	

docker-build: ## Generates the docker image
	docker build -t "newsbot.collector.api" .
	docker image ls | grep newsbot.collector.api
