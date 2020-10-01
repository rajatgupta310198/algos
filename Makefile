GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all: test clean

test:
	sh scripts/run_test.sh

clean:
	$(GOCLEAN) -v


