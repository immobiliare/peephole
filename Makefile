GO ?= go
BINARY_NAME = peephole
GOFLAGS :=
STATIC := 1
DOCKER ?= docker
VERSION := latest
RELEASE ?=
REGISTRY ?=
DOCKER_TAG = $(REGISTRY)peephole:$(VERSION)
LDFLAGS = -X main.version=$(VERSION)

ifeq ($(STATIC), 1)
LDFLAGS += -s -w -extldflags "-static"
endif

ifneq ($(RELEASE),)
VERSION := $(shell git describe --abbrev=0 --tags)
endif

ifneq ($(REGISTRY),)
REGISTRY := $(REGISTRY)/
endif

.PHONY: all

all: build run

clean:
	rm -f $(BINARY_NAME)
	rm -fr vendor/

deps:
	$(GO) mod tidy
	$(GO) mod download

build: deps
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -v -o $(BINARY_NAME) .

docker-build:
	$(DOCKER) build -t $(DOCKER_TAG) .

run: build
	./$(BINARY_NAME) -c ./example.yml

docker-run:
	$(DOCKER) run -v ${PWD}/example.yml:/etc/peephole -p 8080:8080 $(DOCKER_TAG)

docker-push:
	$(DOCKER) push $(DOCKER_TAG)
