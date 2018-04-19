GO_SOURCES := $(shell find src -type f -name "*.go")

default: build

.PHONY: build
build:
	go build -v -o ./bin/sleepuntil $(GO_SOURCES)

.PHONY: lint
lint:
	golint ./src

.PHONY: test
test:
	go test ./src/...

.PHONY: clean
clean:
	rm -rf bin
	mkdir bin