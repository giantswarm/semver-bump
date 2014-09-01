PROJECT=semver-bump

BUILD_PATH := $(shell pwd)/.gobuild

.PHONY=all get-deps build

PROJECT_PATH := "$(BUILD_PATH)/src/github.com/giantswarm"

GOPATH := $(BUILD_PATH)

SOURCE=$(shell find . -name '*.go')

BIN := $(PROJECT)

VERSION := $(shell cat VERSION)

all: .gobuild get-deps $(BIN)

get-deps: .gobuild
	GOPATH=$(GOPATH) go get github.com/coreos/go-semver
	GOPATH=$(GOPATH) go get github.com/spf13/cobra

.gobuild:
	mkdir -p $(PROJECT_PATH)
	cd "$(PROJECT_PATH)" && ln -s ../../../.. $(PROJECT)

$(BIN): $(SOURCE)
	GOPATH=$(GOPATH) go build -a -ldflags "-X main.projectVersion $(VERSION)" -o $(BIN)

clean:
	rm -rf $(BUILD_PATH) $(BIN)
