.PHONY: build
build:
	go build -v ./cmd/app

run:
	go run ./cmd/app/main.go

migrate:
	migrate -path ./db/migration -database 'postgres://postgres@localhost:5432/gotest_db?sslmode=disable' up

drop:
	migrate -path ./db/migration -database 'postgres://postgres@localhost:5432/gotest_db?sslmode=disable' drop

generate_mock_service:
	mockgen -source=pkg/service/service.go -destination=pkg/service/mocks/mock.go
