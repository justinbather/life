BIN_DIR?=./bin
DATABASE_URL?=postgres://postgres:postgres@localhost:5433/life?sslmode=disable

.PHONY: run
run:
	mkdir -p $(BIN_DIR)
	go mod tidy
	migrate -path ./db/migrations/ -database postgres://postgres:postgres@localhost:5433/life?sslmode=disable up
	go build -o $(BIN_DIR)/life-server ./cmd/main.go
	DATABASE_URL=$(DATABASE_URL) ./$(BIN_DIR)/life-server

.PHONY: migrate
migrate: 
	migrate -path ./db/migrations/ -database postgres://postgres:postgres@localhost:5433/life?sslmode=disable up

.PHONY: generate
generate:
	cd ./db && DATABASE_URL=$(DATABASE_URL) sqlc generate
