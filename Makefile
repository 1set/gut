GOCMD=go
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
build:
	time $(GOCMD) build -v ./$(PACKAGE)
test:
	time $(GOTEST) -race -v ./$(PACKAGE)
bench:
	$(GOTEST) -parallel=4 -benchmem -bench=. ./$(PACKAGE)
cover:
	$(GOTEST) -cover -covermode=count ./$(PACKAGE)
all: build test cover bench doc
