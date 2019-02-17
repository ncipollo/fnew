BINARY_NAME=fnew
BUILD_FOLDER=build
LINUX_64_OUTPUT=$(BUILD_FOLDER)/linux64/$(BINARY_NAME)
MAC_64_OUTPUT=$(BUILD_FOLDER)/mac64/$(BINARY_NAME)

GO_BUILD=go build -ldflags "-X main.version=abc" -o


all: deps test build
build: build-linux build-mac
clean:
	go clean
	rm -rf build
deps:
	go get ./...
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) $(LINUX_64_OUTPUT)
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_BUILD) $(MAC_64_OUTPUT)
test:
	go test -v ./...
