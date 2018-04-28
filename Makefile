default: build

.PHONY: build
build:
	cd src/terminalutil; go build -v 
	cd src; go build -v -o ../bin/sleepuntil 

.PHONY: lint
lint:
	golint ./src

.PHONY: test
test: build
	go test ./src/...

.PHONY: clean
clean:
	rm -rf bin
	mkdir bin
