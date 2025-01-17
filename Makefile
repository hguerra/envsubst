.PHONY: build

APP_NAME = envsubst
BUILD_DIR = $(PWD)/build
MAIN_FILE = ./cmd/$(APP_NAME)/main.go

dev_dependencies:
	asdf install
	go get -u github.com/rakyll/gotest
	go get -u github.com/gojekfarm/go-coverage
	asdf reshim golang

dependencies:
	go mod download
	go mod tidy

install: dev_dependencies dependencies

clean:
	rm -rf $(BUILD_DIR)

lint:
	golangci-lint run ./...

test:
	gotest -v -timeout 30s -coverprofile=coverage.out -cover ./...
	go-coverage -f coverage.out --trim

build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

example_stdin:
	go run $(MAIN_FILE) < example-input.tmpl

example_args:
	go run $(MAIN_FILE) -no-empty < example-input.tmpl
