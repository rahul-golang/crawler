
GOCMD=go
GOBUILD=$(GOCMD) build


BINARY_NAME=url-word-count
BUILD_PATH=./bin/$(BINARY_NAME)

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) main.go
   
run:
	$(BUILD_PATH) 