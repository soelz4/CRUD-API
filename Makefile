# Please Don't Change
SRC_DIR := .
.DEFAULT_GOAL := help
BINARY_NAME = main

help:  ## 💬 This Help Message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Linting and Formatting without Fix
lint: ## 🔎 Lint & Format, will not Fix but Sets Exit Code on Error
	gofmt -l $(SRC_DIR) \
	&& gofmt -d main.go \
	&& golangci-lint run main.go

# Linting and Formatting with Try to Fix and Modify Code
lint-fix: ## 📜 Lint & Format, will Try to Fix Errors and Modify Code
	go fmt main.go \
	&& golangci-lint run main.go

# Build Binary File
build: ## 🔨 Build Binary File
	go build -o $(BINARY_NAME) main.go

# RUN
run: build ## 🏃 Run the Web Server Locally at PORT 8080
	$(SRC_DIR)/$(BINARY_NAME)

# Clean up Project
clean: ## 🧹 Clean up Project
	go clean
	# rm $(BINARY_NAME)
