SHELL := /bin/bash -o pipefail
ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

.PHONY : lint

default: lint
	@echo "============= Done ============="

dep:
	@echo "============= Installing dependencies ============="
	go get golang.org/x/tour/gotour
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	@echo "============= Lint project using golangci-linter ============="
	${GOPATH}/bin/golangci-lint run
