BIN_DIR?=./bin
.PHONY: build
build: 
	mkdir -p $(BIN_DIR)
	go mod tidy
	go build -o $(BIN_DIR)/life
