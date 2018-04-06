PROJECT = go-head
PACKAGE = github.com/execjosh/${PROJECT}
TARGET = myhead
EXTLDFLAGS = -extldflags -static $(null)

all: build

.PHONY: test
test: build
	@prove --verbose

.PHONY: build
build: main.go
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o 'bin/${TARGET}' -ldflags '${EXTLDFLAGS}' '${PACKAGE}'
