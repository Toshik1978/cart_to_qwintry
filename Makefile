GOCMD		= go
GOPATH		= $(shell $(GOCMD) env GOPATH)
GOLINT		= golangci-lint
GOCOV		= gocov
GIT_VERSION	= $(shell git rev-list -1 HEAD)

.PHONY: all prereq build lint test mock
.DEFAULT_GOAL := all

all: test build

prereq:
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin
	@$(GOCMD) get github.com/axw/gocov/gocov
	@$(GOCMD) get github.com/golang/mock/gomock
	@$(GOCMD) install github.com/golang/mock/mockgen

mock:
	@$(GOCMD) generate ./...

build:
	@$(GOCMD) build -ldflags "-X main.GitVersion=$(GIT_VERSION)"

lint:
	@$(GOLINT) run ./...

test: lint
	@$(GOCOV) test ./... -v | $(GOCOV) report

test_race: lint
	@$(GOCOV) test ./... -race -v | $(GOCOV) report

clean:
	@$(GOCMD) clean
	@rm -f ./cart_to_qwintry
