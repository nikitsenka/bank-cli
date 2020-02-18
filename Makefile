GOCMD=go
BINARY_NAME=odis

all: clean build

build:
	go build -o bin/$(BINARY_NAME) ./cli/

clean:
	go clean
	rm -rf bin