GO=$(shell which go)
GOBUILD=$(GO) build
GOTEST=$(GO) test

.PHONY:test
test:
	@$(GOTEST) -v

.PHONY:coverage
coverage: test
	@$(GOTEST) -cover

all: build
