.PHONY: default
default: clean fmt build

.PHONY: setup ## Install all the build dependencies
setup:
	@echo Updating dependency tree...
	go mod tidy 
	go mod download
	@echo Updated dependency tree successfully.

.PHONY: clean ## Remove temporary files
clean:
	@go clean -i -testcache .

.PHONY: test ## Run all the tests
test:
	@go test -covermode=atomic -parallel=2 -v -timeout=60s ./...

.PHONY: fmt ## Run gofmt on all go files
fmt:
	@gofmt -w -s -l .

.PHONY: build ## Build a version
build:
	@go build .

.PHONY: run ## Run production
run:
	@./skyfall

.DEFAULT_GOAL := build
