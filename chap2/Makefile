# defines which target is run when no target is specified.
.DEFAULT_GOAL := build

# help make from getting confused when files or dir are named like targets
.PHONY: fmt vet build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean: 
	go clean ./...
