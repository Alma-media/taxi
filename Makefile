GLOLANGCI_LINT_VERSION=1.24.0

install:
	@(cd; GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GLOLANGCI_LINT_VERSION))

lint:
	golangci-lint run --enable-all --disable lll,gochecknoglobals,dupl,interfacer,gochecknoinits,godox,funlen,gocognit,wsl

test:
	go test -v -race ./...

build:
	go build -o bin/taxi main.go

run:
	go run main.go
