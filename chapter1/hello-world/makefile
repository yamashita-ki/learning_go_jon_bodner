.DEFAULT_GOAL := build

fmt:
	go fmt
.PHONY: fmt

lint:
	staticcheck
.PHONY: lint

vet:
	go vet
.PHONY: vet

build:
	go mod tidy
	go build
.PHONY: build