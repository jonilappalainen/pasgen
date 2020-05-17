BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOGET=$(GO) get

BINPATH=main

makedir:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi

build: makedir
	@$(GOBUILD) -o bin

clean:
	@rm -rf $(BUILDPATH)/bin

run: build
	@$(BUILDPATH)/bin/pasgen

all: build
