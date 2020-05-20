BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOGET=$(GO) get
GOTEST=$(GO) test

BINPATH=main

.PHONY:build
build:
	@$(GOBUILD)

.PHONY:test
test: clean build
	@$(GOTEST) -v

.PHONY:coverage
coverage: test
	@$(GOTEST) -cover

all: build
