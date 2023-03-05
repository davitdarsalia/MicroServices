SWAG_DIR := /Users/davitdarsalia/Desktop/microservices/auth/cmd

generate_docs:
	swag init -g $(SWAG_DIR)/main.go -o $(SWAG_DIR)/docs
