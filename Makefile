.PHONY: build run clean dep lint
BINARY_NAME=go.run

build:
	go build -o .build/${BINARY_NAME} cmd/main.go

run: build
	./.build/${BINARY_NAME}

clean:
	go clean
	rm .build/${BINARY_NAME}

dep:
	go mod download
	go mod tidy
	go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@latest
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
lint:
	fieldalignment -fix ./... && golangci-lint run --fix