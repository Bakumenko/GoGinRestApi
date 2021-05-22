.PHONY: build
build:
	go build -v ./cmd

migrate:
	migrate -path ./db/migration -database 'postgres://postgres@localhost:5432/postgres?sslmode=disable' up
.DEFAULT_GOAL := build