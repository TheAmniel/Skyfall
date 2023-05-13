# App info
APP_NAME 			:= skyfall
APP_ENTRY 		:= ./bin/$(APP_NAME)/main.go
CONFIG_NAME 	:= $(APP_NAME).config
CONFIG_FORMAT := toml
VERSION 			:= 1.0.0
COMMIT 				:= $(shell git rev-parse --short HEAD)
BRANCH 				:= $(shell git rev-parse --abbrev-ref HEAD)

GOOS 		:= $(shell go env GOOS)
GOARCH 	:= $(shell go env GOARCH)
COPY 		:= cp

ifeq ($(GOOS), windows)
	TARGET_OS ?= windows
	APP_NAME := $(APP_NAME).exe
else ifeq ($(GOOS), linux)
	TARGET_OS ?= linux
else ifeq ($(GOOS), darwin)
	TARGET_OS ?= darwin
else
	$(error System $(GOOS) is not supported at this time)
endif

# Folders
OUTPUT_FOLDER := .release
BINARY_OUTPUT := $(OUTPUT_FOLDER)/$(APP_NAME)
CONFIG_OUTPUT := $(OUTPUT_FOLDER)/$(CONFIG_NAME).$(CONFIG_FORMAT)

ifeq ($(OS), Windows_NT)
	OUTPUT_FOLDER := .\.release
	CONFIG_OUTPUT := $(OUTPUT_FOLDER)\$(CONFIG_NAME).$(CONFIG_FORMAT)
	COPY 	:= copy
endif

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
	@echo Now building for platform $(GOOS)/$(GOARCH)!
	@go build -ldflags "-s -w \
		-X skyfall/utils.Version=$(VERSION)\
			-X skyfall/utils.Commit=$(COMMIT)\
				-X \"skyfall/utils.BuiltAt=$(shell go run ./bin/built-at/main.go)\"\
					-X skyfall/utils.Branch=$(BRANCH)\
						-X skyfall/utils.AppName=$(APP_NAME)\
							-X skyfall/utils.ConfigFile=$(CONFIG_NAME)"\
				-o $(BINARY_OUTPUT) $(APP_ENTRY) &&\
				$(COPY) .$(CONFIG_NAME).$(CONFIG_FORMAT) $(CONFIG_OUTPUT)
	@echo Successfully built the binary. Use './.release/$(APP_NAME)' to run!

.PHONY: run ## Run production
run:
	@$(BINARY_OUTPUT)

.DEFAULT_GOAL := build
