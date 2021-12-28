BINARY="./base-bin"
MAIN="cmd/base/main.go"
VERSION=1.0.0

default:
	@go build \
	-o ${BINARY} ${MAIN}

.PHONY: default