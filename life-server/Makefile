BIN_DIR?=./bin

.PHONY: run
run:
	mkdir -p $(BIN_DIR)
	go mod tidy
	migrate -path ./db/migrations/ -database postgres://postgres:postgres@localhost:5433/life?sslmode=disable up
	cd ./db && sqlc generate
	go build -o $(BIN_DIR)/life-server ./cmd/main.go
	./$(BIN_DIR)/life-server
