.PHONY: build

version=`git describe --tags`

build: ## Build the application
	@echo "version: ${version}"
	CGO_ENABLED=0 go build -ldflags="-X 'main.Version=${version}'" -o build/commander commander.go

build-all: ## Build the application for supported architectures
	@echo "version: ${version}"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X 'main.Version=${version}'" -o build/${BINARY_NAME}-linux-x86_64 commander.go
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-X 'main.Version=${version}'" -o build/${BINARY_NAME}-linux-aarch64 commander.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-X 'main.Version=${version}'" -o build/${BINARY_NAME}-darwin-x86_64 commander.go
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-X 'main.Version=${version}'" -o build/${BINARY_NAME}-darwin-aarch64 commander.go

install: ## Install the binary
	go get -d ./...
	go get -u golang.org/x/lint/golint

lint: ## Check lint errors
	golint -set_exit_status=1 -min_confidence=1.1 ./...

start: ## Start the application
	go run commander.go