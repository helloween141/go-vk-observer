.DEFAULT_GOAL := build

fmt:
	golangci-lint fmt
.PHONY:fmt

lint:
	golangci-lint run
.PHONY:lint

vet:
	go vet ./...
.PHONY:vet

tidy:
	go mod tidy
.PHONY:tidy

build: tidy fmt lint vet
	docker compose up --build -d
.PHONY:build

up: tidy fmt lint vet
	docker compose up -d
.PHONY:up

restart: down up
.PHONY:restart

down:
	docker compose down
.PHONY:down

shell:
	docker compose exec app sh -c "cd app && sh"
.PHONY:shell