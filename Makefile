.PHONY: default
default: build

include build.properties

GITCOMMIT=$(shell git describe --tags --always --dirty)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%MZ")

APP_SOURCES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: build
build:
	go build -i -v -ldflags="-X '$(importpath)/internal/version.Version=$(version)' -X '$(importpath)/internal/version.GitCommit=$(GITCOMMIT)' -X '$(importpath)/internal/version.BuildDate=$(BUILD_DATE)'" .

.PHONY: codestyle
codestyle:
	gofmt -d -s $(APP_SOURCES)
	go vet
	golint -set_exit_status $(go list ./...)

.PHONY: prepare-tools
prepare-tools:
	go get -u github.com/golang/lint/golint