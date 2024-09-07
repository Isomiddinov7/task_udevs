CURRENT_DIR=$(shell pwd)

APP=$(shell basename "${CURRENT_DIR}")
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest


migration-up:
	migrate -path ./migrations/postgres -database 'postgres://bahodir:1100@0.0.0.0:5432/udevs?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://bahodir:1100@0.0.0.0:5432/udevs?sslmode=disable' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o "${CURRENT_DIR}/bin/${APP}" "${APP_CMD_DIR}/main.go"

run:
	go run cmd/main.go

lint: ## Run golangci-lint with printing to stdout
	golangci-lint -c .golangci.yaml run --build-tags "musl" ./...

swag-init:
	swag init -g api/api.go -o api/docs

