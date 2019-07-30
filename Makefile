SHELL := /bin/bash -o pipefail
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

.PHONY : default
default: lint
	@echo "============= Done ============="

.PHONY: install
install:
	@echo "============= Installing dependencies ============="
	go get golang.org/x/tour/gotour
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	@echo "============= Lint project using golangci-linter ============="
	${GOPATH}/bin/golangci-lint run
