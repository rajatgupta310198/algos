GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all: test build

test:
	$(GOTEST) ./... -v

clean:
	$(GOCLEAN) -v


