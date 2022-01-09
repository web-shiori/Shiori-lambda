.PHONY: install-golangci-lint
install-golangci-lint:
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v1.43.0

.PHONY: test-all
test-all:
	make lint
	make test

.PHONY: test
test:
	go clean ./... && go test ./...

.PHONY: lint
lint:
	make install-golangci-lint
	go clean ./... && golangci-lint cache clean
	golangci-lint run

.PHONY: build
build:
	GOOS=linux go build -o shiori-lambda

.PHONY: zip
zip:
	make build
	zip shiori-lambda.zip shiori-lambda
