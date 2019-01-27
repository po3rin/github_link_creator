.PHONY: dev, prod, build, up, down, test, cover

dev: ## docker-compose build
	cp Dockerfile.dev Dockerfile

prod: ## docker-compose build
	cp Dockerfile.prod Dockerfile

build: ## docker-compose build
	docker-compose build

up: ## docker-compose up -d
	docker-compose up

down: ## docker-compose down
	docker-compose down

test: ## go run test ./...
	go test ./...

cover: ## go test ./... -coverprofile=cover.out
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out
