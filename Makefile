PROJECT=semver-bump

BUILD_PATH := $(shell pwd)/.gobuild

.PHONY=all get-deps build

D0_PATH := "$(BUILD_PATH)/src/github.com/giantswarm"

GOPATH := $(BUILD_PATH)

SOURCE=$(shell find . -name '*.go')

BIN := $(PROJECT)

VERSION := $(shell cat VERSION)

all: .gobuild get-deps $(BIN)

get-deps:
	GOPATH=$(GOPATH) go get github.com/coreos/go-semver
	GOPATH=$(GOPATH) go get github.com/spf13/cobra

.gobuild:
	mkdir -p $(D0_PATH)
	cd "$(D0_PATH)" && ln -s ../../../.. $(PROJECT)

$(BIN): $(SOURCE)
	GOPATH=$(GOPATH) go build -a -ldflags "-X main.version $(VERSION)" -o $(BIN)

clean:
	rm -rf $(BUILD_PATH) $(BIN)
