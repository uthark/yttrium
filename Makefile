.PHONY: default
default: build

include build.properties

GITCOMMIT=$(shell git describe --tags --always --dirty)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%MZ")

APP_SOURCES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOOS=$(shell go env GOOS)

.PHONY: build
build:
	GOOS=$(GOOS) go build -i -v -ldflags="-s -w -X '$(importpath)/internal/version.Version=$(version)' -X '$(importpath)/internal/version.GitCommit=$(GITCOMMIT)' -X '$(importpath)/internal/version.BuildDate=$(BUILD_DATE)'" .

.PHONY: codestyle
codestyle: gofmt golint govet errcheck gocyclo goconst

.PHONY: goconst
goconst:
	goconst $(APP_SOURCES)

.PHONY: gocyclo
gocyclo:
	gocyclo -over 10 $(APP_SOURCES)

.PHONY: govet
govet:
	go vet

.PHONY: gofmt
gofmt:
	gofmt -d -s $(APP_SOURCES)

.PHONY: golint
golint:
	golint -set_exit_status $(shell go list ./...)

.PHONY: errcheck
errcheck:
	errcheck -verbose ./...

.PHONY: prepare-tools
prepare-tools:
	go get -u github.com/golang/lint/golint

.PHONY: set-linux-env
set-linux-env:
	GOOS=linux

.PHONY: linux-build
linux-build: set-linux-env build


.PHONY: docker-build
docker-build: # linux-build
	docker build --pull -t uthark/$(name):$(version)-$(GITCOMMIT) -f build/Dockerfile .



.PHONY: docker-lint
docker-lint:
	docker run -v $(shell pwd)/build/Dockerfile:/Dockerfile:ro redcoolbeans/dockerlint

.PHONY: test
test:
	go test

.PHONY: update-project-dependencies
update-project-dependencies:
	dep ensure -update

.PHONY: all
all:  codestyle build docker-build
