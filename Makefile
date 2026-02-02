APP_NAME := delapore-mouse
CMD_DIR := ./cmd
BUILD_DIR := .

.PHONY: build run clean tidy

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

run: build
	./$(APP_NAME)

clean:
	rm -f $(BUILD_DIR)/$(APP_NAME)

tidy:
	go mod tidy
