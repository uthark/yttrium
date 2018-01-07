default: build

build:
	go build .

codestyle:
	gofmt -d -s .
	go vet