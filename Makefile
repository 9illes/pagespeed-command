BINARY_NAME=pagespeed
BUILD_DIR=build

test: ## Run tests
	go clean -testcache
	go test ./...

coverage: ## Run tests with coverage
	mkdir -p ${BUILD_DIR}
	go test -coverprofile=${BUILD_DIR}/coverage.out ./...
	go tool cover -html=${BUILD_DIR}/coverage.out

build: ## Build the binary file
	go build -o ./${BUILD_DIR}/${BINARY_NAME} main.go

run: ## Run the binary file
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean: ## Remove previous build
	go clean
	rm -f ${BINARY_NAME}
	rm -Rf ./${BUILD_DIR}
