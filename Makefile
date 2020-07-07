OS                   := $(shell uname | tr A-Z a-z )
SHELL                := /bin/bash
PREFIX               := /usr/local
BIN_DIR              := bin

setup:
ifeq ($(shell command -v golint 2> /dev/null),)
	go get -u golang.org/x/lint/golint
endif
ifeq ($(shell command -v goreturns 2> /dev/null),)
	go get -u github.com/sqs/goreturns
endif

.PHONY: fmt
fmt:
	for pkg in $$(go list -f {{.Dir}} ./... | grep -v ^$$(pwd)$$ ); do \
		goreturns -w $$pkg; \
	done

.PHONY: lint
lint:
	for pkg in $$(go list ./...); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

.PHONY: test
test:
	go test $$(go list ./... | tr '\n' ' ')

.PHONY: build
build: .go-set-revision
	$(eval ldflags := -X 'main.revision=$(revision)' -extldflags '-static')
	GOOS=$(OS) GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$(ldflags)" -o $(BIN_DIR)/caff $(BUILD_OPTIONS) main.go

.PHONY: install
install:
ifeq ($(INSTALL_BIN),)
	$(eval bin := caff_$(OS)_amd64)
else
	$(eval bin := $(INSTALL_BIN))
endif
	chmod +x $(BIN_DIR)/caff
	if [ ! -d $(PREFIX)/bin ]; then mkdir -p $(PREFIX)/bin; fi
	cp -a $(BIN_DIR)/caff $(PREFIX)/bin/$(bin)

.PHONY: all
all: setup fmt lint build

.DEFAULT_GOAL := all

.go-set-revision:
	$(eval revision := $(shell if [[ $$REV = "" ]]; then git rev-parse --short HEAD; else echo $$REV;fi;))