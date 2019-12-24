GOCMD=go
GOFMT=goreturns
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODOC=$(GOCMD) doc
GOGET=$(GOCMD) get

ifndef PACKAGE
$(error PACKAGE is not set)
endif

default:
	@echo "build target is required"
	@exit 2
doc:
	$(GODOC) -all ./$(PACKAGE)
fmt:
	$(GOFMT) -l -w ./$(PACKAGE)
build:
	$(GOCMD) build -v ./$(PACKAGE)
test:
	$(GOTEST) -v -race -cover -covermode=atomic -coverprofile=coverage.out ./$(PACKAGE)
	cat coverage.out >> coverage.txt
testdev:
	$(GOTEST) -race -short -cover -covermode=atomic -count 1 ./$(PACKAGE)
bench:
	$(GOTEST) -parallel=4 -run="none" -benchtime="2s" -benchmem -bench=. ./$(PACKAGE)
benchdev:
	$(GOTEST) -parallel=8 -run="none" -benchtime="3s" -benchmem -bench=. ./$(PACKAGE)
all: build test bench doc
