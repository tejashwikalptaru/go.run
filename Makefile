.PHONY: lint deps
BINARY_NAME=go.run

build:
	GOARCH=amd64 GOOS=darwin go build -o .build/${BINARY_NAME} cmd/main.go
 	GOARCH=amd64 GOOS=linux go build -o .build/${BINARY_NAME} cmd/main.go
 	GOARCH=amd64 GOOS=windows go build -o .build/${BINARY_NAME}  cmd/main.go

run: build
	./.build/${BINARY_NAME}

clean:
	go clean
	rm .build/${BINARY_NAME}

dep:
	go mod download
	go mod tidy
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
lint:
	fieldalignment -fix ./... && golangci-lint run --fix