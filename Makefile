.PHONY: all build run gotool clean help

BINARY="MySpace"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go ./config.yaml

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码，并编译生成二进制文件"
	@echo "make build - 编译 Go 代码，生成二进制文件"
	@echo "make run - 直接运行二进制文件"
	@echo "make gotool - 运行 Go工具 'fmt' 和'vet' "