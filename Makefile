APP=consumables

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o ${APP}

.PHONY: run
## run: runs go run
run:
	go run

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@go clean

.PHONY: fmt
## fmt: runs gofmt -w *.go
fmt:
	gofmt -w *.go

.PHONY: help
## help: prints hel message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/	/'