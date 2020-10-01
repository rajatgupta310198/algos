GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all: lint test clean

test:
	sh scripts/run_test.sh

lint:
	go fmt ./...
clean:
	$(GOCLEAN) -v


