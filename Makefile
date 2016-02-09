.DEFAULT_GOAL := build

GOPATH := $(shell echo `pwd`)

dependencies:
	echo "GOPATH="$(GOPATH)
	GOPATH=$(GOPATH) colorgo get ./...

build:
	colorgo build src/main.go
