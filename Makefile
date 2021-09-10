.PHONY: build

build: ## Build the application
	CGO_ENABLED=0 go build -o build/http-commander commander.go

install: ## Install the binary
	go get -d ./...
	go get -u golang.org/x/lint/golint

lint: ## Check lint errors
	golint -set_exit_status=1 -min_confidence=1.1 ./...

start: ## Start the application
	go run commander.go