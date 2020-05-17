BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOGET=$(GO) get
GOTEST=$(GO) test

BINPATH=main

makedir:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/.bin ; fi

build: makedir
	@$(GOBUILD) -o .bin

clean:
	@rm -rf $(BUILDPATH)/.bin

run: build
	@$(BUILDPATH)/.bin/pasgen

test: clean build
	@$(GOTEST) -v

coverage: test
	@$(GOTEST) -cover

all: build
